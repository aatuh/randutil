SHELL := /bin/bash

GO ?= go
GOWORK = off
export GOWORK

GOLANGCI_LINT := github.com/golangci/golangci-lint/v2/cmd/golangci-lint
GOSEC := github.com/securego/gosec/v2/cmd/gosec
GOVULNCHECK := golang.org/x/vuln/cmd/govulncheck
TOOLS_DIR := tools
TOOLS_BIN := $(CURDIR)/.tools/bin
TOOLS_STAMP := $(TOOLS_BIN)/.stamp
GOLANGCI_LINT_BIN := $(TOOLS_BIN)/golangci-lint
GOSEC_BIN := $(TOOLS_BIN)/gosec
GOVULNCHECK_BIN := $(TOOLS_BIN)/govulncheck
FUZZTIME ?= 10s

.PHONY: help test test-ci test-must test-race bench vet lint gosec vuln tidy fmt tools fuzz-smoke clean finalize

help: ## Show help
	@awk 'BEGIN {FS=":.*## "}; \
		/^[a-zA-Z0-9_.-]+:.*## / { \
			if (match($$0, /## .*## /)) { \
				printf "error: multiple ## in help comment for target %s\n", $$1; exit 1; \
			} \
			printf "  %-14s %s\n", $$1, $$2 \
		}' $(MAKEFILE_LIST)

tools: $(TOOLS_STAMP) ## Install pinned Go tools from tools/go.mod

$(TOOLS_STAMP): $(TOOLS_DIR)/go.mod $(TOOLS_DIR)/go.sum
	@mkdir -p $(TOOLS_BIN)
	@GOBIN=$(TOOLS_BIN) $(GO) -C $(TOOLS_DIR) install $(GOLANGCI_LINT)
	@GOBIN=$(TOOLS_BIN) $(GO) -C $(TOOLS_DIR) install $(GOSEC)
	@GOBIN=$(TOOLS_BIN) $(GO) -C $(TOOLS_DIR) install $(GOVULNCHECK)
	@touch $(TOOLS_STAMP)

fmt: ## Run gofmt
	$(GO) fmt ./...

lint: tools ## Run golangci-lint
	$(GOLANGCI_LINT_BIN) run ./...

vuln: tools ## Run govulncheck
	$(GOVULNCHECK_BIN) ./...

gosec: tools ## Run gosec
	$(GOSEC_BIN) ./...

tidy: ## Run go mod tidy
	$(GO) mod tidy
	$(GO) -C $(TOOLS_DIR) mod tidy

test: ## Run unit tests
	$(GO) test ./...

test-ci: ## Run unit tests with randutil_ci build tag
	$(GO) test ./... -tags=randutil_ci

test-must: ## Run unit tests with randutil_must build tag
	$(GO) test ./... -tags=randutil_must

test-race: ## Run unit tests with race detector
	$(GO) test ./... -race -count=1

bench: ## Run benchmarks
	$(GO) test -run=^$$ -bench=. -benchmem ./...

vet: ## Run go vet
	$(GO) vet ./...

fuzz-smoke: ## Run fuzz targets briefly
	FUZZTIME=$(FUZZTIME) scripts/fuzz.sh

clean: ## Clean test cache
	@$(GO) clean -testcache

finalize: ## Run every quality assurance tool
	$(MAKE) tools
	$(MAKE) fmt
	$(MAKE) vet
	$(MAKE) lint
	$(MAKE) vuln
	$(MAKE) gosec
	$(MAKE) tidy
	$(MAKE) test
	$(MAKE) test-ci
	$(MAKE) test-must
	$(MAKE) test-race
	$(MAKE) fuzz-smoke
	$(MAKE) clean
