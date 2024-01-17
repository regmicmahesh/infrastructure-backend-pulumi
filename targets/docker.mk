DOCKER_RUN ?=  \
		docker run \
		--rm \
		-v $(APP_ROOT):/pulumi/projects \
		-w /pulumi/projects \
		--entrypoint make \
		$(BASE_IMAGE_NAME)

docker/pulumi-login:
	@$(DOCKER_RUN) pulumi-login

docker/pulumi-preview:
	@$(DOCKER_RUN) pulumi-preview
	
docker/pulumi-up:
	@$(DOCKER_RUN) pulumi-up

docker/pulumi-destroy:
	@$(DOCKER_RUN) pulumi-destroy

docker/go-vendor:
	@$(DOCKER_RUN) go-vendor

