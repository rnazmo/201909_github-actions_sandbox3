

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	# go get -u golang.org/x/lint/golint
	go vet ./...
	golint -set_exit_status ./...
