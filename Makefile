EXECUTABLE=jrctl
VERSION=$(shell git describe --tags `git rev-list --tags --max-count=1`)
OS=linux
ARCH=amd64

.PHONY: help bump build-all clean docs format package

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

bump: ## Bump version in source files based on latest git tag
	VERSION=$(VERSION); sed -E -i '' "s/(Version-)([0-9.]+)(-green)/\1$$VERSION\3/g" README.md
	VERSION=$(VERSION); sed -E -i '' "s/(version: )([0-9.]+)/\1$$VERSION/g" nfpm.yaml
	VERSION=$(VERSION); sed -E -i '' "s/(const VersionString string = \")([0-9.]+)(\")/\1$$VERSION\3/g" sdk/version/version.go

build: ## Build for specified platforms
	env GOOS=$(OS) GOARCH=$(ARCH) go build -v -o "bin/$(EXECUTABLE)_$(OS)_$(ARCH)" -ldflags="-s -w" -trimpath ./cmd/jrctl/main.go

build-all: ## Build for all platforms
	make OS=darwin ARCH=amd64 build
	make OS=linux  ARCH=amd64 build
	make OS=linux  ARCH=arm64 build

clean: ## Delete built binaries
	mkdir -p dist bin
	rm -rf dist/* bin/* nfpm.generated.yaml

docs: ## Generate documentation
	mkdir -p docs man
	rm -rf man/* docs/*.md
	DOCS=true NO_COLOR=true go run tools/generate-docs.go

format: ## Format code with goimports
	gofmt -w -s cmd internal tools sdk pkg
	goimports -w cmd internal tools sdk pkg

package: clean format build-all docs ## Package binary for many distributions
	GOARCH=amd64 envsubst < nfpm.yaml > nfpm.generated.yaml
	GOARCH=amd64 nfpm pkg --config nfpm.generated.yaml --packager deb --target dist/$(EXECUTABLE)_$(VERSION)_linux_amd64.deb
	GOARCH=amd64 nfpm pkg --config nfpm.generated.yaml --packager rpm --target dist/$(EXECUTABLE)_$(VERSION)_linux_amd64.rpm
	GOARCH=arm64 envsubst < nfpm.yaml > nfpm.generated.yaml
	GOARCH=arm64 nfpm pkg --config nfpm.generated.yaml --packager deb --target dist/$(EXECUTABLE)_$(VERSION)_linux_arm64.deb
	GOARCH=arm64 nfpm pkg --config nfpm.generated.yaml --packager rpm --target dist/$(EXECUTABLE)_$(VERSION)_linux_arm64.rpm
	tar -czvf ./dist/$(EXECUTABLE)_$(VERSION)_darwin_amd64.tar.gz man bin/$(EXECUTABLE)_darwin_amd64
	rm -f nfpm.generated.yaml
