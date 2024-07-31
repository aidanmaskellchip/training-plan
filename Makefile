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
	make docker-volume-create
	make docker-network-create
	make docker-build-go-image
	make go-mod-tidy
	make go-vendor-download
	make -i stop-db
	make setup-db
	make run-migrations
	make start-server

docker-build-go-image:
	docker build --target go --file ./Dockerfile --tag ${DOCKER_GO_IMAGE_REF} .

docker-network-create:
	docker network inspect ${DOCKER_NETWORK_DB_REF} >/dev/null  2>&1 || docker network create ${DOCKER_NETWORK_DB_REF}

docker-volume-create:
	docker volume create --driver local --opt type=nfs --opt o=addr=host.docker.internal,rw,nolock,hard,nointr,nfsvers=3 --opt device=:${PWD} --name=${DOCKER_VOLUME_APP_REF}

setup-db:
	docker run -d --name training-plan_db -p 8432:5432 --network ${DOCKER_NETWORK_DB_REF} -e POSTGRES_PASSWORD=training-plan_password -e POSTGRES_USER=training-plan_user -e POSTGRES_DB=training-plan_db -e PGDATA="/var/lib/postgresql/data/pgdata" postgres:latest

stop-db:
	docker stop training-plan_db
	docker container remove training-plan_db

start-server:
	go run ./cmd/api

run-migrations:
	go run cmd/fixture/migrate/main.go

##############
##### Go #####
##############
go-run-cmd:
	docker run -it --rm --env-file ./.env -v ${DOCKER_VOLUME_APP_REF}:/go/src/app --network ${DOCKER_NETWORK_DB_REF} ${DOCKER_GO_IMAGE_REF} ${cmd}

go-get-lib:
	make go-run-cmd cmd='go get -v ${lib}'

go-mod-tidy:
	make go-run-cmd cmd='go mod tidy'

go-vendor-download:
	make go-run-cmd cmd='go mod vendor'

go-tests:
	go test -v ./...