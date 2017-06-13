SHELL = bash
MAKEFLAGS += --silent
GOFILES ?= $(shell go list ./... | grep -v /vendor/)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

.PHONY: build

build:
	go fmt && go clean && go build

test:
	go test $(GOFILES)

update:
	curl https://glide.sh/get | sh
	glide up && glide install

cover:
	go test $(GOFILES) --cover
