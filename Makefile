.PHONY: test clean build-testchild

clean:
	go clean -cache -testcache

build:
	go build ./...

build-testchild:
	go build -o ./pipes/testchild ./cmd/testchild

test: clean build-testchild
	go test -race -count=1 ./...
