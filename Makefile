BIN_PATH = ./bin

GOLANGCI_LINT_VERSION = v1.31.0
GOLANGCI_LINT_BIN = $(BIN_PATH)/golangci-lint
GOLANGCI_LINT_INSTALL_URL = https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh

clean:
	@go clean
	@rm -rf $(BIN_PATH)
.PHONY: clean

test:
	@go test -race -v ./...
.PHONY: test

lint: install-golangci-lint
	@$(GOLANGCI_LINT_BIN) run --color always ./...
.PHONY: lint

install-golangci-lint:
	@test -f $(GOLANGCI_LINT_BIN) || \
		{ curl -sfL $(GOLANGCI_LINT_INSTALL_URL) | sh -s -- -b $(BIN_PATH) $(GOLANGCI_LINT_VERSION) ; }
.PHONY: install-golangci-lint
