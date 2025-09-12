# Makefile for Infra Dashboard API

.PHONY: help build run test clean docker-build docker-run docker-compose-up docker-compose-down deps lint fmt

# Default target
help:
	@echo "Available targets:"
	@echo "  build              - Build the application binary"
	@echo "  run                - Run the application locally"
	@echo "  test               - Run tests"
	@echo "  clean              - Clean build artifacts"
	@echo "  deps               - Download and tidy dependencies"
	@echo "  lint               - Run golangci-lint"
	@echo "  fmt                - Format Go code"
	@echo "  docker-build       - Build Docker image"
	@echo "  docker-run         - Run Docker container"
	@echo "  docker-compose-up  - Start all services with docker-compose"
	@echo "  docker-compose-down- Stop all services with docker-compose"

# Application name and version
APP_NAME := infra-dashboard
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME := $(shell date +%Y-%m-%dT%H:%M:%S%z)

# Build flags
LDFLAGS := -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)"

# Build the application
build:
	@echo "Building $(APP_NAME)..."
	go build $(LDFLAGS) -o $(APP_NAME) cmd/main.go

# Run the application locally
run:
	@echo "Running $(APP_NAME)..."
	go run cmd/main.go

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -f $(APP_NAME)
	go clean

# Download and tidy dependencies
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

# Run golangci-lint (requires golangci-lint to be installed)
lint:
	@echo "Running linter..."
	golangci-lint run

# Format Go code
fmt:
	@echo "Formatting code..."
	go fmt ./...
	goimports -w .

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t $(APP_NAME):$(VERSION) .
	docker tag $(APP_NAME):$(VERSION) $(APP_NAME):latest

# Run Docker container
docker-run:
	@echo "Running Docker container..."
	docker run --rm -p 8080:8080 \
		-e DB_HOST=host.docker.internal \
		-e DB_PORT=5432 \
		-e DB_USER=postgres \
		-e DB_PASSWORD=postgres \
		-e DB_NAME=infra_dashboard \
		-e DB_SSLMODE=disable \
		-e SERVER_PORT=8080 \
		$(APP_NAME):latest

# Start all services with docker-compose
docker-compose-up:
	@echo "Starting services with docker-compose..."
	docker-compose up -d --build

# Stop all services with docker-compose
docker-compose-down:
	@echo "Stopping services with docker-compose..."
	docker-compose down

# Development setup
dev-setup: deps
	@echo "Setting up development environment..."
	@if [ ! -f .env ]; then \
		echo "Creating .env file from .env.example..."; \
		cp .env.example .env; \
	fi
	@echo "Development setup complete!"
	@echo "1. Make sure PostgreSQL is running"
	@echo "2. Update .env with your database credentials"
	@echo "3. Run 'make run' to start the application"

# Database operations (requires psql)
db-create:
	@echo "Creating database..."
	createdb infra_dashboard

db-drop:
	@echo "Dropping database..."
	dropdb infra_dashboard

db-reset: db-drop db-create
	@echo "Database reset complete"

# Install development tools
install-tools:
	@echo "Installing development tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest

# Run the application with hot reload (requires air: go install github.com/cosmtrek/air@latest)
dev:
	@echo "Starting development server with hot reload..."
	air
