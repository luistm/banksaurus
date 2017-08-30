install: deps unit-tests integration-tests
	$(shell go install)

unit-tests:
	go test ./... -v -short

coverage-unit:
	go test ./... -v -short -cover

coverage-package:
	go test $(PACKAGE) -v -short -cover -covermode=count -coverprofile=count.out
	go tool cover -func=count.out

integration-tests:
	go test ./... -v

deps:
	$(shell glide install)

