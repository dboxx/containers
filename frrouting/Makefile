TAG   ?= v8.0.1
NAME  ?= frrouting
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
	@echo " > docker run --name ${NAME} -d --rm --net=host --pid=host --privileged --mount type=bind,source=/lib/modules,target=/lib/modules ${IMAGE}"
	@echo " > docker run --name ${NAME} -d --rm --cap-add NET_ADMIN --cap-add NET_BROADCAST -p 179:179 --mount type=bind,source=/lib/modules,target=/lib/modules ${IMAGE}"
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

.PHONY: latest
latest: ## mark image as latest
	@docker tag ${IMAGE}:${TAG} ${IMAGE}:latest
	@docker push ${IMAGE}:latest

.PHONY: run
run:   ## run image
	@docker run -d --rm --net=host --pid=host --privileged --name ${NAME} --mount type=bind,source=/lib/modules,target=/lib/modules ${IMAGE}:${TAG}
