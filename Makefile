install: deps
	go install

build: clean deps
	go build -o go-bank-cli

clean:
	- rm go-bank-cli

test: unit-tests system-tests

unit-tests:
	go test ./... -v -short

coverage-unit:
	go test ./... -short -cover

system-tests: build
	go test -run=^TestSystem$

coverage-package:
	go test $(PACKAGE)  -short -cover -covermode=count -coverprofile=count.out
	go tool cover -func=count.out
	
deps:
	glide install

style:
	golint
	go fmt ./...
