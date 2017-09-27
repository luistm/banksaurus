install: deps
	go install ./cmd/bank

build: clean
	go build -i -o bank ./cmd/bank 

clean:
	- rm bank

test: unit-tests system-tests coverage-unit

unit-tests:
	go test ./... -short

coverage-unit:
	go test ./... -short -cover

system-tests: build
	go test ./... -run System -v

coverage-package:
	go test $(PACKAGE)  -short -cover -covermode=count -coverprofile=count.out
	go tool cover -func=count.out
	
deps:
	glide install

style:
	golint
	go fmt ./...
