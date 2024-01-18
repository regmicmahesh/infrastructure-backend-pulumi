DOCKER_RUN ?=  \
		docker run \
		-it \
		--rm \
		-v $(APP_ROOT):/pulumi/projects \
		-v $(APP_ROOT)/.cache/:/root/.cache \
		-w /pulumi/projects \
		-e PULUMI_STACK_NAME=$(PULUMI_STACK_NAME) \
		-e LAYER_NAME=$(LAYER_NAME) \
		--entrypoint make \
		$(BASE_IMAGE_NAME)


shell:
	@/bin/bash

docker/shell:
	@$(DOCKER_RUN) shell

docker/pulumi-login:
	@$(DOCKER_RUN) pulumi-login

docker/pulumi-init:
	@$(DOCKER_RUN) pulumi-init

docker/pulumi-preview:
	@$(DOCKER_RUN) pulumi-preview
	
docker/pulumi-up:
	@$(DOCKER_RUN) pulumi-up

docker/pulumi-destroy:
	@$(DOCKER_RUN) pulumi-destroy

docker/pulumi-cancel:
	@$(DOCKER_RUN) pulumi-cancel

docker/pulumi-output:
	@$(DOCKER_RUN) pulumi-output

docker/pulumi-refresh:
	@$(DOCKER_RUN) pulumi-refresh

docker/pulumi-graph:
	@$(DOCKER_RUN) pulumi-graph 

docker/pulumi-stack-export:
	@$(DOCKER_RUN) pulumi-stack-export

docker/pulumi-stack:
	@$(DOCKER_RUN) pulumi-stack

docker/go-vendor:
	@$(DOCKER_RUN) go-vendor

