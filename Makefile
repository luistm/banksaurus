install: deps unit-tests integration-tests
	$(shell go install)

unit-tests:
	go test ./... -v -short

coverage-unit:
	go test ./... -v -short -cover

integration-tests:
	go test ./... -v

deps:
	$(shell glide install)

