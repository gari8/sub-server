.PHONY: cross
cross: main.go  ## Build binaries for cross platform.
	mkdir -p pkg
	@# darwin
	@for arch in "amd64" "386"; do \
		GOOS=darwin GOARCH=$${arch} make build; \
		zip pkg/kube-prompt_$(VERSION)_darwin_$${arch}.zip kube-prompt; \
	done;
	@# linux
	@for arch in "amd64" "386" "arm64"; do \
		GOOS=linux GOARCH=$${arch} make build; \
		zip pkg/kube-prompt_$(VERSION)_linux_$${arch}.zip kube-prompt; \
	done;