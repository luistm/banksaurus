unit-tests:
	go test ./... -v -short

integration-test:
	go test ./... -v

deps:
	$(shell glide install)

install: deps unit-tests integration-tests
	$(shell go install)
