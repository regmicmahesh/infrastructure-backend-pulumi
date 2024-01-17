
NON_INTERACTIVE_FLAGS := --non-interactive --color=never

pulumi-login:
	@pulumi login s3://$(PULUMI_BACKEND_STATE_BUCKET_NAME) --non-interactive --color=never
	@pulumi stack select $(PULUMI_STACK_NAME)

pulumi-preview: pulumi-login
	@pulumi preview $(NON_INTERACTIVE_FLAGS)
	
pulumi-up: pulumi-login
	@pulumi up --skip-preview $(NON_INTERACTIVE_FLAGS)

pulumi-destroy: pulumi-login
	@pulumi destroy --skip-preview $(NON_INTERACTIVE_FLAGS)
