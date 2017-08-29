unit-tests:
	go test ./... -v -short

integration-tests:
	go test ./... -v

deps:
	$(shell glide install)

install: deps unit-tests integration-tests
	$(shell go install)
