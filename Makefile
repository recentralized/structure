default: build lint test test-examples

build:
	go build ./...

lint:
	golint ./...

test:
	go test ./...

test-examples: test-example-index test-example-meta

test-example-%:
	@./examples/check $*

