SHELL := /bin/bash

GO ?= go
GOWORK ?= off

TOOLS := golangci-lint gosec govulncheck
GOLANGCI_LINT_VERSION ?= v1.64.8
GOSEC_VERSION ?= v2.22.11
GOVULNCHECK_VERSION ?= v1.1.4

.PHONY: help test test-race lint gosec vuln tidy fmt tools clean finalize

help: ## Show help
	@awk 'BEGIN {FS=":.*## "}; \
		/^[a-zA-Z0-9_.-]+:.*## / { \
			if (match($$0, /## .*## /)) { \
				printf "error: multiple ## in help comment for target %s\n", $$1; exit 1; \
			} \
			printf "  %-14s %s\n", $$1, $$2 \
		}' $(MAKEFILE_LIST)

tools: ## Install lint/vuln tools
	@$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)
	@$(GO) install github.com/securego/gosec/v2/cmd/gosec@$(GOSEC_VERSION)
	@$(GO) install golang.org/x/vuln/cmd/govulncheck@$(GOVULNCHECK_VERSION)

fmt: ## Run gofmt
	@GOWORK=$(GOWORK) $(GO) fmt ./...

lint: tools ## Run golangci-lint
	@GOWORK=$(GOWORK) golangci-lint run ./...

vuln: tools ## Run govulncheck
	@GOWORK=$(GOWORK) govulncheck ./...

gosec: tools ## Run gosec
	@GOWORK=$(GOWORK) gosec ./...

tidy: ## Run go mod tidy
	@GOWORK=$(GOWORK) $(GO) mod tidy

test: ## Run unit tests
	@GOWORK=$(GOWORK) $(GO) test ./...

test-race: ## Run unit tests with race detector
	@GOWORK=$(GOWORK) $(GO) test ./... -race -count=1

clean: ## Clean test cache
	@$(GO) clean -testcache

finalize: ## Run every quality assurance tool
	$(MAKE) tools
	$(MAKE) fmt
	$(MAKE) lint
	$(MAKE) vuln
	$(MAKE) gosec
	$(MAKE) tidy
	$(MAKE) test
	$(MAKE) test-race
	$(MAKE) clean
