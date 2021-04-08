NAME  = asterisk
ORG   = dboxx
TAG   = v16.2.1
IMAGE = ${ORG}/${NAME}

# Colors
COLNO   = \033[0m
COLERR  = \033[31;01m
COLWARN = \033[33;01m
COLMENU = \033[36m

.PHONY: help
help:   ## show list
	@echo "Git: https://github.com/dboxx/containers"
	@echo "Hub: https://hub.docker.com/r/${IMAGE}"
	@echo " > docker pull ${IMAGE}"
	@echo "-----"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "${COLMENU}%-10s${COLNO} %s\n", $$1, $$2}'

.PHONY: pull
pull:    ## pull image
	@docker pull ${IMAGE}:${TAG}

.PHONY: push
push:    ## push image
	@docker push ${IMAGE}:${TAG}

.PHONY: build
build:   ## build new image
	@docker build . -t ${IMAGE}:${TAG}

.PHONY: latest
latest:  ## mark image as latest
	@docker tag ${IMAGE}:${TAG} ${IMAGE}:latest
	@docker push ${IMAGE}:latest

.PHONY: run
run:    ## run image
	@docker run --name ${NAME} --net=host --pid=host --privileged --rm -it ${IMAGE}:${TAG}

.PHONY: start
start:   ## start container
	@docker run --name ${NAME} --privileged -d --rm ${IMAGE}:${TAG}

.PHONY: stop
stop:    ## stop container
	@docker stop ${NAME}

.PHONY: exec
exec:    ## exec command
	@docker exec -it ${NAME} bash

.PHONY: status
status:  ## show container status
	@docker ps -af name=${NAME}

.PHONY: clean
clean:   ## remove image
	@docker rmi ${IMAGE}:${TAG}
