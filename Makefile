VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' \
           -X 'main.revision=$(REVISION)'

.PHONY: build
build: main.go  ## Build a binary.
	$(GO) build -ldflags "$(LDFLAGS)"

.PHONY: cross
cross: main.go  ## Build binaries for cross platform.
	mkdir -p pkg
	@# darwin
	@for arch in "amd64" "386"; do \
		GOOS=darwin GOARCH=$${arch} make build; \
		zip pkg/sub-server_$(VERSION)_darwin_$${arch}.zip sub-server; \
	done;
	@# linux
	@for arch in "amd64" "386" "arm64"; do \
		GOOS=linux GOARCH=$${arch} make build; \
		zip pkg/sub-server_$(VERSION)_linux_$${arch}.zip sub-server; \
	done;