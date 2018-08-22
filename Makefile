.DEFAULT_GOAL := help
.PHONY: all

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

local: deps ## Installs banksaurus locally
	go install -ldflags '-s -w' ./cmd/bscli

build: clean ## Builds the project
	go build -i -o bscli ./cmd/bscli

clean: ## Cleans binary created by make build
	- rm banksaurus

tests:  unit-tests integration-tests acceptance-tests ## Runs all tests

unit-tests: ## Runs unit tests
	go test ./... -run Unit

integration-tests: ## Runs integration tests
	go test ./... -run Integration

coverage-unit: ## Runs coverage for unit tests
	go test ./... -run Unit -cover

acceptance-tests: build ## Runs acceptance tests
	- go test ./... -run Acceptance
	- rm bscli

coverage-package: ## Runs coverage for package $PACKAGE
	go test $(PACKAGE)  -run Unit -cover -covermode=count -coverprofile=count.out
	go tool cover -func=count.out
	
deps: ## Ensures dependencies are met
	dep ensure

style: ## Formats the project
	goimports -w .
	go fmt ./...
