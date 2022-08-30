NAME  = go-hello
ORG   = dboxx
TAG   = v1.0
REPO  = docker.io
IMAGE = ${ORG}/${NAME}

# Container Runtimes: [docker|podman]
PODMAN_VERSION = $(shell podman --version 2>/dev/null)
DOCKER_VERSION = $(shell docker --version 2>/dev/null)

ifdef PODMAN_VERSION
    RUNTIME = podman
else ifdef DOCKER_VERSION
    RUNTIME = docker
else
    $(error [ ERROR ] container runtime can't be detected)
endif

# Colors
COLNO   = \033[0m
COLERR  = \033[31;01m
COLWARN = \033[33;01m
COLMENU = \033[36m

.PHONY: help
help:   ## show list
	@echo "Git: https://github.com/dboxx/containers"
	@echo "Hub: https://hub.docker.com/r/${IMAGE}"
	@echo " > ${RUNTIME} pull ${IMAGE}"
	@echo "-----"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "${COLMENU}%-10s${COLNO} %s\n", $$1, $$2}'

.PHONY: pull
pull:   ## pull image
	${RUNTIME} pull ${REPO}/${IMAGE}:${TAG}

.PHONY: push
push:   ## push image
	${RUNTIME} push ${REPO}/${IMAGE}:${TAG}

.PHONY: run
run:    ## run image
	${RUNTIME} run --name ${NAME} --rm -it ${REPO}/${IMAGE}:${TAG} /app/${NAME}

.PHONY: build
build:  ## build new image
	${RUNTIME} build . -t ${REPO}/${IMAGE}:${TAG}

.PHONY: latest
latest: ## mark image as latest
	${RUNTIME} tag ${REPO}/${IMAGE}:${TAG} ${REPO}/${IMAGE}:latest
	${RUNTIME} push ${REPO}/${IMAGE}:latest

.PHONY: clean
clean:  ## remove image
	${RUNTIME} rmi ${REPO}/${IMAGE}:${TAG}

