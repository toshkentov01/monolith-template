CURRENT_DIR=$(shell pwd)
APP=$(shell basename ${CURRENT_DIR})

APP_CMD_DIR=${CURRENT_DIR}/cmd
PKG_LIST := $(shell go list ./... | grep -v /vendor/)

IMG_NAME=${APP}
REGISTRY=${REGISTRY:-861701250313.dkr.ecr.us-east-1.amazonaws.com}
TAG=latest
ENV_TAG=latest

ifneq (,$(wildcard ./.env))
	include .env
endif

ifdef CI_COMMIT_BRANCH
        include .build_info
endif

make create-env:
	cp ./.env.example ./.env

set-env:
	./scripts/set-env.sh ${CURRENT_DIR}

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

proto-gen:
	./scripts/gen-proto.sh  ${CURRENT_DIR}

migrate-jeyran: set-env
	env POSTGRES_HOST=${POSTGRES_HOST} env POSTGRES_PORT=${POSTGRES_PORT} env POSTGRES_USER=${POSTGRES_USER} env POSTGRES_PASSWORD=${POSTGRES_PASSWORD} env POSTGRES_DB=${POSTGRES_DB} ./scripts/migrate-jeyran.sh

pull-proto-module:
	git submodule update --init --recursive

update-proto-module:
	git submodule update --remote --rebase

clear:
	rm -rf ${CURRENT_DIR}/bin/*

network:
	docker network create --driver=bridge ${NETWORK_NAME}

mark-as-staging-image:
	docker tag ${REGISTRY}/${IMG_NAME}:${TAG} ${REGISTRY}/${IMG_NAME}:staging
	docker push ${REGISTRY}/${IMG_NAME}:staging

mark-as-production-image:
	docker tag ${REGISTRY}/${IMG_NAME}:${TAG} ${REGISTRY}/${IMG_NAME}:production
	docker push ${REGISTRY}/${IMG_NAME}:production

build-image:
	docker build --rm -t ${REGISTRY}/${IMG_NAME}:${TAG} .

push-image:
	docker push ${REGISTRY}/${IMG_NAME}:${TAG}

run-dev: set-env
	go run cmd/main.go

dep:
	go get -v -d ./...

lint:
	golint -set_exit_status ${PKG_LIST}

unit-tests: set-env ## Run unit-tests
	go test -mod=vendor -v -cover -short ${PKG_LIST}

race: set-env ## Run data race detector
	go test -mod=vendor -race -short ${PKG_LIST}

msan: set-env ## Run memory sanitizer. If this test fails, you need to write the following command: export CC=clang (if you have installed clang)
	env CC=clang env CXX=clang++ go test -mod=vendor -msan -short ${PKG_LIST}

delete-branches:
	${CURRENT_DIR}/scripts/delete-branches.sh

swag-gen:
	echo ${REGISTRY}
	swag init -g cmd/main.go -o api/docs

run-redis:
	docker run --rm -p 6379:6379 --name redis-local --detach redis:4.0.9

stop-redis:
	docker stop redis-local

update-go-deps:
	@echo ">> updating Go dependencies"
	@for m in $$(go list -mod=readonly -m -f '{{ if and (not .Indirect) (not .Main)}}{{.Path}}{{end}}' all); do \
		go get $$m; \
	done
	go mod tidy


docker.run: docker.fiber

docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network

docker.fiber.build:
	docker build -t fiber .

docker.fiber: docker.fiber.build
	docker run \
		--publish 8000:8000 \
		fiber

docker.stop: docker.stop.fiber

docker.stop.fiber:
	docker stop cgapp-fiber

ifneq (,$(wildcard vendor))
	go mod vendor
endif

.PHONY: vendor


