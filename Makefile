EXECUTABLE=jrctl
VERSION=$(shell git describe --tags `git rev-list --tags --max-count=1`)
LINUX=$(EXECUTABLE)_linux_amd64
DARWIN=$(EXECUTABLE)_darwin_amd64

.PHONY: help clean docs

build: linux darwin ## Build for all platforms
	@echo version: $(VERSION)

linux: $(LINUX)

darwin: $(DARWIN)

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -i -v -o "bin/$(LINUX)" -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/jrctl/main.go

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -i -v -o "bin/$(DARWIN)" -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/jrctl/main.go

clean: ## Delete built binaries
	rm -f "bin/$(LINUX)" "bin/$(DARWIN)"

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

docs: ## Generate documentation
	mkdir -p docs man
	rm -rf man/* docs/*.md
	JR_DOCS=true JR_COLOR=false go run tools/generate-docs.go

package: build ## Package binary for many distributions
	mkdir -p ./dist
	rm -f ./dist/*
	nfpm pkg --packager deb --target ./dist
	nfpm pkg --packager rpm --target ./dist
	tar -czvf ./dist/$(EXECUTABLE)-$(VERSION)-darwin.tar.gz man bin/$(DARWIN)
