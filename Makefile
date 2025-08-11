DOCKER_GO_IMAGE_REF=training-plan_go
DOCKER_GO_CMD_IMAGE_REF=golang-alpine:latest
UP_ARGS=-d
DEPLOY_ARGS=--verbose --debug *
DOCKER_VOLUME_APP_REF=training-plan_app
DOCKER_NETWORK_DB_REF=training-plan_db
DOCKER_VOLUME_DB_REF=training-plan_db

# Colours used in help
RED      := $(shell tput -Txterm setaf 1)
GREEN    := $(shell tput -Txterm setaf 2)
WHITE    := $(shell tput -Txterm setaf 7)
CYAN     := $(shell tput -Txterm setaf 6)
YELLOW   := $(shell tput -Txterm setaf 3)
RESET    := $(shell tput -Txterm sgr0)

#################################################
################# DOCKER ########################
#################################################
setup:
	make docker-build-cmd
	make docker-network-create
	make docker-build-app
	make go-vendor-download
	make docker-up
	make run-migrations

docker-build-cmd:
	DOCKER_NETWORK_DB_REF=${DOCKER_NETWORK_DB_REF} docker build --target builder --tag ${DOCKER_GO_CMD_IMAGE_REF} .

docker-network-create:
	docker network inspect ${DOCKER_NETWORK_DB_REF} >/dev/null  2>&1 || docker network create ${DOCKER_NETWORK_DB_REF}

docker-build-app:
	DOCKER_NETWORK_DB_REF=${DOCKER_NETWORK_DB_REF} docker build --target runner --tag ${DOCKER_GO_IMAGE_REF} .

docker-up:
	DOCKER_NETWORK_DB_REF=${DOCKER_NETWORK_DB_REF} docker-compose -f docker-compose.yml up --build --force-recreate ${UP_ARGS}

run-migrations:
	make go-run-cmd cmd='go run cmd/fixture/migrate/main.go'

##############
##### Go #####
##############
go-run-cmd:
	docker run --rm --env-file ./.env -v ${PWD}:/app --network ${DOCKER_NETWORK_DB_REF} ${DOCKER_GO_CMD_IMAGE_REF} ${cmd}

go-tests:
	make go-run-cmd cmd='go test -v ./...'

go-get-lib:
	make go-run-cmd cmd='go get -v ${lib}'

go-mod-tidy:
	make go-run-cmd cmd='go mod tidy'

go-vendor-download:
	make go-run-cmd cmd='go mod vendor'

go-lint:
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.55.2 golangci-lint run -v
