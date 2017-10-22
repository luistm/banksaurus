install: deps
	go install ./cmd/bankcli

build: clean
	go build -i -o bankcli ./cmd/bankcli

clean:
	- rm bankcli

test: unit-tests system-tests coverage-unit

unit-tests:
	go test ./... -run Unit

coverage-unit:
	go test ./... -run Unit -cover

system-tests: build
	- go test ./... -run System
	- rm bankcli

coverage-package:
	go test $(PACKAGE)  -run Unit -cover -covermode=count -coverprofile=count.out
	go tool cover -func=count.out
	
deps:
	dep ensure

style:
	golint
	go fmt ./...
