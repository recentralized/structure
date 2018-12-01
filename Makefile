default: build lint test test-examples

build:
	go build ./...

lint:
	golint ./...
	gofmt -s -d $$(find . -name '*.go')

test:
	go test ./...

gofmt:
	gofmt -s -w $$(find . -name '*.go')

cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

test-examples: test-example-index test-example-meta

test-example-%:
	@./examples/check $*

