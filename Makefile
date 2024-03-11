.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

test: fmt
	go test . -v
.PHONY:test

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet test
	go build .
.PHONY:build