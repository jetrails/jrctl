EXECUTABLE=jrctl
OUTPUT=bin
LINUX=$(EXECUTABLE)_linux_amd64
DARWIN=$(EXECUTABLE)_darwin_amd64
VERSION=$(shell git describe --tags --always --long --dirty)

.PHONY: help clean docs

build: linux darwin ## Build for all platforms
	@echo version: $(VERSION)

linux: $(LINUX)

darwin: $(DARWIN)

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -i -v -o "$(OUTPUT)/$(LINUX)" -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/jrctl/main.go

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -i -v -o "$(OUTPUT)/$(DARWIN)" -ldflags="-s -w -X main.version=$(VERSION)"  ./cmd/jrctl/main.go

clean: ## Delete built binaries
	rm -f "$(OUTPUT)/$(LINUX)" "$(OUTPUT)/$(DARWIN)"

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

docs: ## Generate documentation
	go run tools/generate-docs.go
