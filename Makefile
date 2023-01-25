# This should match the version defined in .github/workflows/lint.yaml
WANTED_LINT_VERSION := 1.50.1
APP_NAME := go-vbo365-mon
PACKAGE_NAME := github.com/patrikcze/go-vbo365
LINT_VERSION := $(shell golangci-lint version | cut -d' ' -f4)
HAS_LINT := $(shell which golangci-lint)

#
INSTALL_LINT_PAGE := "https://golangci-lint.run/usage/install/"
BAD_LINT_MSG := "Missing golangci-lint version $(WANTED_LINT_VERSION). Visit $(INSTALL_LINT_PAGE) for instructions on how to install"

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
IMAGE_NAME=go-vbo365-mon
CONTAINER_NAME=go-vbo365-mon
REGISTRY=docker.io

all: build

build:
	@echo "Building $(APP_NAME)..."
	$(GOBUILD) -o $(APP_NAME) $(PACKAGE_NAME)

test:
	@echo "Running tests for $(APP_NAME)..."
	$(GOTEST) -v $(PACKAGE_NAME)/...

lint:
	@echo "Linting $(APP_NAME)..."
	golangci-lint run
	staticcheck $(PACKAGE_NAME)/...
	revive -config revive.toml $(PACKAGE_NAME)/...

dep:
	@echo "Installing dependencies for $(APP_NAME)..."
	dep ensure
	$(GOGET) -u honnef.co/go/tools/cmd/staticcheck
	$(GOGET) -u github.com/mgechev/revive

clean:
	@echo "Cleaning $(APP_NAME)..."
	$(GOCLEAN) $(PACKAGE_NAME)
	rm $(APP_NAME)
