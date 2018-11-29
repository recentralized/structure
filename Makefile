default: build lint test

build: 
	go build ./...

lint:
	golint ./...

test:
	go test ./...
