TAG   ?= v1.2.4
NAME  ?= quagga
IMAGE ?= dboxx/${NAME}

# Colors
COLNO=\033[0m
COLERR=\033[31;01m
COLWARN=\033[33;01m
COLMENU=\033[36m

.PHONY: help
help:   ## show help
	@echo "Git: https://github.com/dboxx/containers"
	@echo "Hub: https://hub.docker.com/r/${IMAGE}"
	@echo " > docker pull ${IMAGE}"
	@echo " > docker run --name ${NAME} --rm -it --net=host --pid=host --privileged ${IMAGE}"
	@echo " > docker run --name ${NAME} --rm -it --cap-add NET_ADMIN --cap-add NET_BROADCAST -p 2605:2605 -p 179:179 ${IMAGE}"
	@echo "-----"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' ${MAKEFILE_LIST} | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "${COLMENU}%-10s${COLNO} %s\n", $$1, $$2}'

.PHONY: pull
pull:   ## pull image
	@docker pull ${IMAGE}:${TAG}

.PHONY: push
push:   ## push image
	@docker push ${IMAGE}:${TAG}

.PHONY: build
build:   ## build image
	@docker build . -t ${IMAGE}:${TAG} -f build/Dockerfile \
		&& echo "Building ... [ ${COLOK}OK${COLNO} ]" \
		|| echo "Building ... [ ${COLERR}Error${COLNO} ]"

.PHONY: run
run:   ## run image
	@docker run --rm --name vtysh -it ${IMAGE}:${TAG} vtysh

.PHONY: latest
latest: ## mark image as latest
	@docker tag ${IMAGE}:${TAG} ${IMAGE}:latest
	@docker push ${IMAGE}:latest

.PHONY: local-up
local-up:   ## run app locally
	@docker-compose -f build/docker-compose.yaml up -d --remove-orphans \
		&& echo "Local upping ... [ ${COLOK}OK${COLNO} ]" \
		|| echo "Local upping ... [ ${COLERR}Error${COLNO} ]"

.PHONY: local-down
local-down:   ## stop local app
	@docker-compose -f build/docker-compose.yaml down \
		&& echo "Local downing ... [ ${COLOK}OK${COLNO} ]" \
		|| echo "Local downing ... [ ${COLERR}Error${COLNO} ]"

.PHONY: local-ps
local-ps:   ## check app locally
	@docker-compose -f build/docker-compose.yaml ps

.PHONY: local-sh
local-sh:   ## run vtysh
	@docker exec -it quagga vtysh
