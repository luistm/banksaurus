install: deps unit-tests integration-tests
	$(shell go install)

unit-tests:
	go test ./... -v -short

coverage-unit:
	go test ./... -short -cover

coverage-package:
	go test $(PACKAGE)  -short -cover -covermode=count -coverprofile=count.out
	go tool cover -func=count.out

integration-tests:
	go test ./... -v

deps:
	$(shell glide install)

style:
	golint
	go fmt ./...