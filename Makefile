GOLANGCI_VERSION := $(shell golangci-lint --version 2>/dev/null)
DB_DOCKER_PATH := ../5e-database

.PHONY: help ## Show this help.
help:
	@grep -he '^.PHONY:.*##' $(MAKEFILE_LIST) | sed -e 's/ *##/:\t/' | sed -e 's/^.PHONY: *//'

.PHONY: build ## Build app and output to bin/main.
build:
	go build -o bin/main cmd/biscuit-dnd-go/main.go

.PHONY: run ## Run app.
run:
	go run cmd/biscuit-dnd-go/main.go

.PHONY: test ## Run all tests.
test:
	go test ./...

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

.PHONY: builddb ## Build the docker database image. Pass the path to your local copy of https://github.com/5e-bits/5e-database as DB_DOCKER_PATH=<path>, or default ../5e-database will be used.
builddb:
	docker build -t 5e-database $(DB_DOCKER_PATH)

.PHONY: startdb ## Start the database. Pass the path to your local copy of https://github.com/5e-bits/5e-database as DB_DOCKER_PATH=<path>, or default ../5e-database will be used.
startdb: builddb
	docker run -p 27017:27017 -t 5e-database:latest


