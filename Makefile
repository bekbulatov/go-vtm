GO := go
pkgs = $(shell $(GO) list ./... | grep -v /vendor/)

all: format vet build test

format:
	@echo ">> formatting code"
	@$(GO) fmt $(pkgs)

vet:
	@echo ">> vetting code"
	@$(GO) vet $(pkgs)

test:
	@echo ">> running tests"
	@$(GO) test -short -coverprofile cover.out
	@$(GO) tool cover -html=cover.out -o cover.html

build:
	@echo ">> building binaries"
	@$(GO) build

lint:
	@echo ">> linting code"
	@golint $(pkgs)

examples:
	make -C examples all

.PHONY: all format vet lint test coverage build compile examples
