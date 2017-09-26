install: deps
	go install ./cmd/bank

build: clean deps
	go build -o bank ./cmd/bank 

clean:
	- rm bank

test: unit-tests system-tests

unit-tests:
	go test ./... -short

coverage-unit:
	go test ./... -short -cover

system-tests: build
	go test ./... -run=^TestSystem$ 

coverage-package:
	go test $(PACKAGE)  -short -cover -covermode=count -coverprofile=count.out
	go tool cover -func=count.out
	
deps:
	glide install

style:
	golint
	go fmt ./...
