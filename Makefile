PLATFORM=local
RELEASE_VERSION=1.0.0
DOCKER_TAG_ROOT=stock
DOCKER_TAG_NAME=datainjestor

all: \
	build

.PHONY: build
build: \
	build-docker

.PHONY: build-docker
build-docker: 
	@docker build . --target app \
	--tag ${DOCKER_TAG_ROOT}/${DOCKER_TAG_NAME}:${RELEASE_VERSION} \
	--tag ${DOCKER_TAG_ROOT}/${DOCKER_TAG_NAME}:latest \
	--platform ${PLATFORM}

.PHONY: 
test: \
	unit-test

.PHONY: unit-test
unit-test:
	@docker build . --target unit-test