install: deps
	go install

build: clean deps
	go build -o go-cli-bank

clean:
	rm go-cli-bank

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
