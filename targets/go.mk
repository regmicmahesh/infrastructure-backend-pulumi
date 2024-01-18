go-vendor:
	@go mod vendor

go-build:
	@echo Building...
	@go build -o . $(APP_ROOT)/layers/$(LAYER_NAME)
