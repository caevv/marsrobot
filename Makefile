SHELL=/bin/bash

.PHONY: all
all: deps lint test

.PHONY: deps
deps:
	go mod download

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	golangci-lint run
