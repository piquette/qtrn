
# Copyright © 2018 Piquette Capital, LLC
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
QTRN_VERSION = 0.5.1
# Linker flags
LDFLAGS := -X "github.com/piquette/qtrn/version.Version=$(QTRN_VERSION)"

.PHONY: release
release:


.PHONY: test
test:
	@go test -v ./...

.PHONY: vet
vet:

.PHONY: dev
dev:
	go build -v -ldflags="$(LDFLAGS)"
	cp qtrn $(GOPATH)/bin/
