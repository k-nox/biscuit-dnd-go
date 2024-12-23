GOLANGCI_VERSION := $(shell golangci-lint --version 2>/dev/null)
DB_DOCKER_PATH := ../5e-database

.PHONY: help ## Show this help.
help:
	@grep -he '^.PHONY:.*##' $(MAKEFILE_LIST) | sed -e 's/ *##/:\t/' | sed -e 's/^.PHONY: *//'

.PHONY: install ## Install dependencies.
install:
	go get -t ./...
	go mod tidy

.PHONY: build ## Build app and output to bin/main.
build: install
	go build -o bin/main cmd/biscuit-dnd-go/main.go

.PHONY: run ## Run app.
run: install
	go run cmd/biscuit-dnd-go/main.go

.PHONY: test ## Run all tests.
test: install
	go test -v ./...

.PHONY: test-unit ## Only run unit tests that don't require MongoDB.
test-unit: install
	go test -v -short ./...

.PHONY: lint ## Run golangci-lint, checking if golangci-lint is installed and prompting to install via homebrew if necessary.
lint: check-linter-installed
	golangci-lint run

.PHONY: lint-fix ## Run golangci-lint with --fix, checking if golangci-lint is installed and prompting to install via homebrew if necessary.
lint-fix: check-linter-installed
	golangci-lint run --fix

.PHONY: check-linter-installed ## Check if golangci-lint is installed and optionally install via homebrew if needed.
check-linter-installed:
ifndef GOLANGCI_VERSION
	@echo "Golangci-lint not installed, do you want to install it with homebrew? [y/N] " && read ans && [ $${ans:-N} = y ]
	brew install golangci-lint
endif

.PHONY: db ## Build & run the docker database image. To use a local copy of 5e-databases, uncomment the `build` key and comment out the `image` key in docker-compose.yaml.
db-up:
	docker compose up --detach --build db

db-down:
	docker compose down --volumes --remove-orphans