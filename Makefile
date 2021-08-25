EXECUTABLE=jrctl
VERSION=$(shell git describe --tags `git rev-list --tags --max-count=1`)
LINUX=$(EXECUTABLE)_linux_amd64
DARWIN=$(EXECUTABLE)_darwin_amd64

.PHONY: help clean docs format

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

bump: ## Bump version in source files based on latest git tag
	VERSION=$(VERSION); sed -E -i '' "s/(Version-)([\d.]+)(-green)/\1$$VERSION\3/g" ./README.md
	VERSION=$(VERSION); sed -E -i '' "s/(jrctl_)([\d.]+)(_)/\1$$VERSION\3/g" ./README.md
	VERSION=$(VERSION); sed -E -i '' "s/(jrctl-)([\d.]+)(\.)/\1$$VERSION\3/g" ./README.md
	VERSION=$(VERSION); sed -E -i '' "s/(download\/)([\d.]+)(\/jrctl)/\1$$VERSION\3/g" ./README.md
	VERSION=$(VERSION); sed -E -i '' "s/(version: )([\d.]+)/\1$$VERSION/g" ./nfpm.yaml
	VERSION=$(VERSION); sed -E -i '' "s/(const VersionString string = \")([\d.]+)(\")/\1$$VERSION\3/g" ./sdk/version/version.go

build: bump linux darwin ## Build for all platforms
	@echo version: $(VERSION)

linux: $(LINUX)

darwin: $(DARWIN)

$(LINUX):
	env GOOS=linux GOARCH=amd64 go build -v -o "bin/$(LINUX)" -ldflags="-s -w -X main.version=$(VERSION)" -trimpath ./cmd/jrctl/main.go

$(DARWIN):
	env GOOS=darwin GOARCH=amd64 go build -v -o "bin/$(DARWIN)" -ldflags="-s -w -X main.version=$(VERSION)" -trimpath ./cmd/jrctl/main.go

clean: ## Delete built binaries
	rm -f "bin/$(LINUX)" "bin/$(DARWIN)"
	rm -rf ./dist/*

docs: ## Generate documentation
	mkdir -p docs man
	rm -rf man/* docs/*.md
	DOCS=true NO_COLOR=true go run tools/generate-docs.go

format: ## Format code with goimports
	gofmt -w -s cmd internal tools sdk pkg
	goimports -w cmd internal tools sdk pkg

package: format build docs ## Package binary for many distributions
	mkdir -p ./dist
	rm -f ./dist/*
	nfpm pkg --packager deb --target ./dist
	nfpm pkg --packager rpm --target ./dist
	tar -czvf ./dist/$(EXECUTABLE)-$(VERSION)-darwin.tar.gz man bin/$(DARWIN)
