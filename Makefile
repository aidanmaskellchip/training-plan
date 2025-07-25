DOCKER_GO_IMAGE_REF=training-plan_go
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
#	make docker-volume-create
	make docker-network-create
	make docker-build-go-image
	make docker-up
	make go-mod-tidy
	make go-vendor-download
	make run-migrations
	make start-server

docker-build-go-image:
	docker build --target go --file ./Dockerfile --tag ${DOCKER_GO_IMAGE_REF} .

docker-network-create:
	docker network inspect ${DOCKER_NETWORK_DB_REF} >/dev/null  2>&1 || docker network create ${DOCKER_NETWORK_DB_REF}

docker-volume-create:
	docker volume create --driver local --opt type=nfs --opt o=addr=host.docker.internal,rw,nolock,hard,nointr,nfsvers=3 --opt device=:${PWD} --name=${DOCKER_VOLUME_APP_REF}

docker-up:
	DOCKER_NETWORK_DB_REF=${DOCKER_NETWORK_DB_REF} docker-compose -f docker-compose.yml up --build --force-recreate ${UP_ARGS}

start-server:
	make go-run-cmd cmd='go run ./cmd/api'

run-migrations:
	go run cmd/fixture/migrate/main.go

##############
##### Go #####
##############
go-run-cmd:
	docker run --rm --env-file ./.env -v ${DOCKER_VOLUME_APP_REF}:/go/src/app --network ${DOCKER_NETWORK_DB_REF} ${DOCKER_GO_IMAGE_REF} ${cmd}

go-tests:
	go test -v ./...

go-get-lib:
	make go-run-cmd cmd='go get -v ${lib}'

go-mod-tidy:
	make go-run-cmd cmd='go mod tidy'

go-build-cmd:
	make go-run-cmd cmd='env GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o .aws-sam/build/${buildir}/bootstrap cmd/lambda/${handler}/main.go'

go-vendor-download:
	make go-run-cmd cmd='go mod vendor'