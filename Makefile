GO := go
pkgs = $(shell $(GO) list ./... | grep -v /vendor/)

all: format build vet test examples

format:
	@echo ">> formatting code"
	@$(GO) fmt $(pkgs)

build:
	@echo ">> building binaries"
	@$(GO) build

vet:
	@echo ">> vetting code"
	@$(GO) vet $(pkgs)

test:
	@echo ">> running tests"
	@$(GO) test -short -coverprofile cover.out
	@$(GO) tool cover -html=cover.out -o cover.html

examples:
	@echo ">> building examples"
	make -C examples all

lint:
	@echo ">> linting code"
	@golint $(pkgs)

.PHONY: all format build vet test examples lint
