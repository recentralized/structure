default: build lint test test-examples

build:
	go build ./...

lint:
	golint ./...

test:
	go test ./...

test-examples:
	find ./examples -name "main.go" | xargs -n1 -IX go run X > /dev/null

