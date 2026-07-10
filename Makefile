# ==============================================================================
# Goflix - Development Makefile
# ==============================================================================

.DEFAULT_GOAL := help

GREEN  := \033[0;32m
YELLOW := \033[1;33m
BLUE   := \033[0;34m
RED    := \033[0;31m
RESET  := \033[0m

.PHONY: help run build \
        fmt vet test check \
        tidy update \
        clean doctor

help: ## Show available commands
	@echo ""
	@echo "$(BLUE)Goflix Development Commands$(RESET)"
	@echo ""
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z0-9_-]+:.*##/ {printf "  \033[32m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""

# ==============================================================================
# Development
# ==============================================================================

run: ## Run the server
	go run ./cmd/goflix

build: ## Build local binary
	@mkdir -p bin
	go build -o bin/goflix ./cmd/goflix
	@echo "$(GREEN)✓ Binary generated in bin/goflix$(RESET)"

# ==============================================================================
# Quality
# ==============================================================================

fmt: ## Format source code
	go fmt ./...

vet: ## Run go vet
	go vet ./...

test: ## Run unit tests
	go test ./...

check: fmt vet test ## Run all quality checks

# ==============================================================================
# Dependencies
# ==============================================================================

tidy: ## Clean go.mod / go.sum
	go mod tidy

update: ## Update dependencies
	go get -u ./...
	go mod tidy

# ==============================================================================
# Utilities
# ==============================================================================

clean: ## Remove generated files
	rm -rf bin
	rm -f goflix.db

doctor: ## Display development environment
	@echo ""
	@echo "$(BLUE)Environment$(RESET)"
	@echo ""
	@go version
	@git --version
