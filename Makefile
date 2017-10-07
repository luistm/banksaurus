install: deps
	go install ./cmd/bank

build: clean
	go build -i -o bank ./cmd/bank 

clean:
	- rm bank

test: unit-tests system-tests coverage-unit

unit-tests:
	go test ./... -run Unit

coverage-unit:
	go test ./... -run Unit -cover

system-tests: build
	go test ./... -run System -v

coverage-package:
	go test $(PACKAGE)  -run Unit -cover -covermode=count -coverprofile=count.out
	go tool cover -func=count.out
	
deps:
	dep ensure

style:
	golint
	go fmt ./...
