.PHONY: build clean test

all: build

build:
	go build -v .

clean:
	go clean -i .

test:
	go test ./...