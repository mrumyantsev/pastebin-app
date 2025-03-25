.DEFAULT_GOAL := build
.SILENT:

include ./.env

export ALPINE_VERSION
export POSTGRES_VERSION
export POSTGRES_DATABASE_LOCAL_DIR
export MONGO_DATABASE_LOCAL_DIR
export MONGO_VERSION
export GOLANG_VERSION
export SERVER_APP_NAME
export NGINX_VERSION
export WEB_LOCAL_DIR
export WEB_PORT
export APP_START_TIMEOUT_SECONDS
export APP_STOP_TIMEOUT_SECONDS
export LOGGER_GLOBAL_LEVEL
export DATABASE_MIGRATE
export POSTGRES_DATABASE_HOST
export POSTGRES_DATABASE_PORT
export POSTGRES_DATABASE_USERNAME
export POSTGRES_DATABASE_PASSWORD
export POSTGRES_DATABASE_NAME
export POSTGRES_DATABASE_MAX_CONNS
export POSTGRES_DATABASE_MAX_CONN_IDLE_TIME_MINUTES
export MONGO_DATABASE_HOST
export MONGO_DATABASE_PORT
export MONGO_DATABASE_USERNAME
export MONGO_DATABASE_PASSWORD
export MONGO_DATABASE_NAME
export SERVER_HOST
export SERVER_PORT
export SERVER_READ_TIMEOUT_SECONDS
export SERVER_WRITE_TIMEOUT_SECONDS
export SERVER_MAX_HEADER_BYTES
export SERVER_ENABLE_DEBUG_MODE

.PHONY: build
build:
	docker compose build

.PHONY: run
run:
	docker compose up

.PHONY: build-postgres-container
build-postgres-container:
	docker run \
	--name pb-postgres \
	-d \
	-e POSTGRES_USER=${POSTGRES_DATABASE_USERNAME} \
	-e POSTGRES_PASSWORD=${POSTGRES_DATABASE_PASSWORD} \
	-e POSTGRES_DB=${POSTGRES_DATABASE_NAME} \
	-e PGDATA=/data \
	-v ${POSTGRES_DATABASE_LOCAL_DIR}:/data \
	-p ${POSTGRES_DATABASE_PORT}:5432 \
	postgres:${POSTGRES_VERSION}-alpine${ALPINE_VERSION}

.PHONY: build-mongo-container
build-mongo-container:
	docker run \
	--name pb-mongo \
	-d \
	-e MONGO_INITDB_ROOT_USERNAME=${MONGO_DATABASE_USERNAME} \
	-e MONGO_INITDB_ROOT_PASSWORD=${MONGO_DATABASE_PASSWORD} \
	-e MONGO_INITDB_DATABASE=${MONGO_DATABASE_NAME} \
	-v ${MONGO_DATABASE_LOCAL_DIR}:/data/db \
	-p ${MONGO_DATABASE_PORT}:27017 \
	mongo:${MONGO_VERSION}-rc1-jammy

.PHONY: run-postgres-container
run-postgres-container:
	docker start pb-postgres

.PHONY: run-mongo-container
run-mongo-container:
	docker start pb-mongo

.PHONY: build-local
build:
	go build -o ./build/${SERVER_APP_NAME} ./cmd/${SERVER_APP_NAME}

.PHONY: run-local
run:
	./build/${SERVER_APP_NAME}

.PHONY: run-fast
run-fast:
	go run ./cmd/${SERVER_APP_NAME}

.PHONY: test
test:
	go test -v ./...
