.DEFAULT_GOAL := build

.PHONY: clean fmt vet build

build_dir := build

clean:
	go clean -x
	rm $(build_dir)/* 
fmt:
	go fmt ./...

vet: fmt
	go vet -fix ./...

build: vet
	go build -o $(build_dir)/rev
