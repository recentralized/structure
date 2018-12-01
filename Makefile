default: build lint test test-examples

build:
	go build ./...

lint:
	golint ./...

test:
	go test ./...

cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

test-examples: test-example-index test-example-meta

test-example-%:
	@./examples/check $*

