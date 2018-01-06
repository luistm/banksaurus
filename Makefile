.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

local: deps ## Installs bankcli locally
	go install -ldflags '-s -w' ./cmd/bankcli

build: clean ## Builds the project
	go build -i -o bankcli ./cmd/bankcli

clean: ## Cleans binary created by make build
	- rm bankcli

test: unit-tests system-tests coverage-unit ## Runs all tests

unit-tests: ## Runs unit tests
	go test ./... -run Unit

coverage-unit: ## Runs coverage for unit tests
	go test ./... -run Unit -cover

system-tests: build ## Runs system tests
	- go test ./... -run System
	- rm bankcli

coverage-package: ## Runs coverage for package $PACKAGE
	go test $(PACKAGE)  -run Unit -cover -covermode=count -coverprofile=count.out
	go tool cover -func=count.out
	
deps: ## Ensures dependencies are met
	dep ensure

style: ## Formats the project
	golint
	go fmt ./...
