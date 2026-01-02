.PHONY: help run dev build test clean deps

APP_NAME = go-crud
BINARY_NAME = $(APP_NAME)

help: ## Display this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

run: ## Run the application
	@echo "Starting $(APP_NAME)..."
	@go run main.go

dev: ## Run with hot reload (development mode)
	@echo "Starting $(APP_NAME) in development mode with hot reload..."
	@go run github.com/githubnemo/CompileDaemon -command="./$(BINARY_NAME)"

build: ## Build the application
	@echo "Building $(APP_NAME)..."
	@go build -o $(BINARY_NAME) .

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

clean: ## Remove binary and cache
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@go clean -cache

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download

lint: ## Run linter
	@echo "Running golangci-lint..."
	@golangci-lint run

.PHONY: help
.DEFAULT_GOAL := help