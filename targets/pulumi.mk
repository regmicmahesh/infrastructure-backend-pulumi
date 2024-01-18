
export FLAGS = --non-interactive --color=never --cwd=$(APP_ROOT)/layers/$(LAYER_NAME)

pulumi-login:
	@pulumi login s3://$(PULUMI_BACKEND_STATE_BUCKET_NAME) --non-interactive --color=never
	@pulumi stack select $(PULUMI_STACK_NAME) $(FLAGS)

pulumi-init:
	@pulumi login s3://$(PULUMI_BACKEND_STATE_BUCKET_NAME) --non-interactive --color=never
	@pulumi stack init $(PULUMI_STACK_NAME) $(FLAGS)

pulumi-preview: pulumi-login go-build
	@pulumi preview --diff --show-replacement-steps $(FLAGS)
	
pulumi-up: pulumi-login go-build
	@pulumi up --skip-preview $(FLAGS)

pulumi-destroy: pulumi-login go-build
	@pulumi destroy --skip-preview $(FLAGS)

pulumi-cancel: pulumi-login go-build
	@pulumi cancel $(FLAGS)

pulumi-output: pulumi-login go-build
	@pulumi stack output --show-secrets $(FLAGS)

pulumi-refresh: pulumi-login go-build
	@pulumi refresh --skip-preview $(FLAGS) 

pulumi-graph: pulumi-login
	@pulumi stack graph ./stack-graph --short-node-name --ignore-parent-edges

pulumi-stack-export: pulumi-login
	@pulumi stack export --file stack-output.json

pulumi-stack: pulumi-login
	@pulumi stack --show-urns
