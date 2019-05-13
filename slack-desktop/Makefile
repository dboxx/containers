TAG = v3.4.0
NAME = slack-desktop
IMAGE = dboxx/$(NAME)

# Colors
COLNO=\033[0m
COLERR=\033[31;01m
COLWARN=\033[33;01m
COLMENU=\033[36m

.PHONY: help pull push run build latest

help:   ## show list
	@echo "Git: https://github.com/dboxx/containers"
	@echo "Hub: https://hub.docker.com/r/$(IMAGE)"
	@echo "-----"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "$(COLMENU)%-10s$(COLNO) %s\n", $$1, $$2}'

pull:   ## pull image
	@docker pull $(IMAGE):$(TAG)

push:   ## push image
	@docker push $(IMAGE):$(TAG)

run:    ## run image
	@docker run --rm --name $(NAME) \
		          --network host \
		          -e DISPLAY=$(DISPLAY) \
							-v /tmp/.X11-unix/X0:/tmp/.X11-unix/X0 \
							-v $(HOME)/.Xauthority:/root/.Xauthority \
							-it $(IMAGE):$(TAG) slack

build:  ## build new image
	@docker build . -t $(IMAGE):$(TAG)

latest: ## mark image as latest
	@docker tag $(IMAGE):$(TAG) $(IMAGE):latest
	@docker push $(IMAGE):latest
