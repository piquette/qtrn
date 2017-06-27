
# Copyright 2017 FlashBoys All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# qtrn version
QTRN_VERSION = 0.3

# Go and compilation related variables
BUILD_DIR ?= out

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

VETREPORT = vet.report
PACKAGES := go list ./... | grep -v /vendor
SOURCE_DIRS = cli version

# Linker flags
LDFLAGS := -X "github.com/FlashBoys/qtrn/version.Version=$(QTRN_VERSION)"


# Build targets

vendor:
	glide install -v

$(BUILD_DIR)/$(GOOS)-$(GOARCH):
	mkdir -p $(BUILD_DIR)/$(GOOS)-$(GOARCH)

$(BUILD_DIR)/darwin-amd64/qtrn: vendor $(BUILD_DIR)/$(GOOS)-$(GOARCH)
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build --installsuffix cgo -ldflags="$(LDFLAGS)" -o $(BUILD_DIR)/darwin-amd64/qtrn

$(BUILD_DIR)/linux-amd64/qtrn: vendor $(BUILD_DIR)/$(GOOS)-$(GOARCH)
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build --installsuffix cgo -ldflags="$(LDFLAGS)" -o $(BUILD_DIR)/linux-amd64/qtrn

$(BUILD_DIR)/windows-amd64/qtrn.exe: vendor $(BUILD_DIR)/$(GOOS)-$(GOARCH)
	CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build --installsuffix cgo -ldflags="$(LDFLAGS)" -o $(BUILD_DIR)/windows-amd64/qtrn.exe

$(GOPATH)/bin/gh-release:
	go get -u github.com/progrium/gh-release/...

.PHONY: release
release: clean fmtcheck test $(GOPATH)/bin/gh-release cross
	mkdir -p release
	gnutar -zcf release/qtrn-$(QTRN_VERSION)-darwin-amd64.tgz LICENSE README.md -C $(BUILD_DIR)/darwin-amd64 qtrn
	gnutar -zcf release/qtrn-$(QTRN_VERSION)-linux-amd64.tgz LICENSE README.md -C $(BUILD_DIR)/linux-amd64 qtrn
	zip -j release/qtrn-$(QTRN_VERSION)-windows-amd64.zip LICENSE README.md $(BUILD_DIR)/windows-amd64/qtrn.exe
	gh-release checksums sha256
	gh-release create FlashBoys/qtrn $(QTRN_VERSION) master v$(QTRN_VERSION)

.PHONY: cross
cross: $(BUILD_DIR)/darwin-amd64/qtrn $(BUILD_DIR)/linux-amd64/qtrn $(BUILD_DIR)/windows-amd64/qtrn.exe

.PHONY: clean
clean:
	go clean -v -i ./...
	rm -rf $(VETREPORT)
	rm -rf $(BUILD_DIR)
	rm -rf release

.PHONY: test
test:
	@go test -v $(shell $(PACKAGES))

.PHONY: fmtcheck
fmtcheck:
	@gofmt -l -s $(SOURCE_DIRS) | grep ".*\.go"; if [ "$$?" = "0" ]; then exit 1; fi


.PHONY: vet
vet:
	@go vet $(PACKAGES) > $(VETREPORT)

.PHONY: dev
dev: clean
	go build -v -ldflags="$(LDFLAGS)"
	cp qtrn $(GOPATH)/bin/

.PHONY: up
up:
	curl https://glide.sh/get | sh
	glide up
	glide install
