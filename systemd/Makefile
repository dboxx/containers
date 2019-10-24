TAG = v241.7
NAME = systemd
IMAGE = dboxx/$(NAME)

# Colors
COLNO=\033[0m
COLERR=\033[31;01m
COLWARN=\033[33;01m
COLMENU=\033[36m

.PHONY: help pull push build latest start stop exec status clean

help:   ## show list
	@echo "Git: https://github.com/dboxx/containers"
	@echo "Hub: https://hub.docker.com/r/$(IMAGE)"
	@echo " > docker pull $(IMAGE)"
	@echo "-----"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "$(COLMENU)%-10s$(COLNO) %s\n", $$1, $$2}'

pull:    ## pull image
	@docker pull $(IMAGE):$(TAG)

push:    ## push image
	@docker push $(IMAGE):$(TAG)

build:   ## build new image
	@docker build . -t $(IMAGE):$(TAG)

latest:  ## mark image as latest
	@docker tag $(IMAGE):$(TAG) $(IMAGE):latest
	@docker push $(IMAGE):latest

start:   ## start container
	@docker run --name $(NAME) --privileged -d --rm $(IMAGE):$(TAG)

stop:    ## stop container
	@docker stop $(NAME)

exec:    ## exec command
	@docker exec -it $(NAME) bash

status:  ## show container status
	@docker ps -af name=$(NAME)

clean:   ## remove image
	@docker rmi $(IMAGE):$(TAG)
