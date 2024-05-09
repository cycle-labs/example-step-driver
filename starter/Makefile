## ----------------------------------------------------------------------
## Makefile for Example Step Driver
## Preqrequisites: GOPATH env var is set to GOPATH=$HOME/go
## ----------------------------------------------------------------------

.PHONY: help
help:     ## Show this help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: install_oapi_codegen
install_oapi_codegen: ## Install the oapi-codegen tool
	go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest

.PHONY: gen_api_server_model
gen_api_server_model:  ## Generate the types from the OpenAPI spec for the server targer
	$$GOPATH/bin/oapi-codegen -generate types -package api openapi/v1/schema.openapi.yaml > ./api/models.generated.go

.PHONY: gen_api_gin
gen_api_gin: ## Generate the Gin boilerplate from the OpenAPI spec
	$$GOPATH/bin/oapi-codegen -generate gin -package api openapi/v1/schema.openapi.yaml > ./api/gin.generated.go

.PHONY: gen_api_server_all
gen_api_server_all: gen_api_server_model gen_api_gin ## Generate types and server boilerplate from the OpenAPI spec

.PHONY: gen_api_client_model
gen_api_client_model:  ## Generate the types from the OpenAPI spec for the client target
	$$GOPATH/bin/oapi-codegen -generate types -package api openapi/v1/schema.openapi.yaml > ./api/models.generated.go

.PHONY: gen_api_client
gen_api_client: ## Generate the Gin boilerplate from the OpenAPI spec
	$$GOPATH/bin/oapi-codegen -generate client -package api openapi/v1/schema.openapi.yaml > ./api/client.generated.go

.PHONY: gen_api_client_all
gen_api_client_all: gen_api_client_model gen_api_client ## Generate types and server boilerplate from the OpenAPI spec

.PHONY: air
air: ## Start live reloads with Air
	$$GOPATH/bin/air
