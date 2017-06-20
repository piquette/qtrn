SHELL = bash
PROJECT = qtrn
VETREPORT = vet.report
GOFILES ?= $(shell go list ./... | grep -v /vendor/)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
GOPATH=$(shell go env GOPATH)
VERSION=0.2

LDFLAGS += -X "github.com/FlashBoys/qtrn/version.Version=$(VERSION)"
LDFLAGS += -X "main.CommitHash=$(shell git rev-parse HEAD)"
MAKEFLAGS += --silent

.PHONY: dev clean test up vet release

dev: clean
	echo "Building.."
	go build -v -ldflags '$(LDFLAGS)'
	cp $(PROJECT) $(GOPATH)/bin/

clean:
	echo "Cleaning.."
	go clean -v -i ./...
	rm -rf $(VETREPORT)
	rm -rf builds/

test:
	go test -cover $(GOFILES)

up:
	curl https://glide.sh/get | sh
	glide up
	glide install

vet:
	go vet $(GOFILES) > $(VETREPORT)

release: clean test linux windows darwin

linux:
	echo "Building for linux.."
	GOOS=linux GOARCH=amd64 go build -v -ldflags '$(LDFLAGS)' -o ./builds/$(PROJECT)-linux-$(GOARCH)

darwin:
	echo "Building for mac.."
	GOOS=darwin GOARCH=amd64 go build -v -ldflags '$(LDFLAGS)' -o ./builds/$(PROJECT)-darwin-$(GOARCH)

windows:
	echo "Building for windows.."
	GOOS=windows GOARCH=amd64 go build -ldflags '$(LDFLAGS)' -o ./builds/$(PROJECT)-windows-$(GOARCH).exe
