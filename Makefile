unit-tests:
	go test ./... -v -short

integration-test:
	go test ./... -v

deps:
	$(shell glide install)
