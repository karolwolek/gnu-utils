.DEFAULT_GOAL := build

.PHONY: clean fmt vet build

clean:
	go clean -x

fmt:
	go fmt ./...

vet: fmt
	go vet -fix ./...

build: vet
	go build
