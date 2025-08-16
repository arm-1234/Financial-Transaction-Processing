.PHONY: help build run test clean docker-up docker-down migrate-up migrate-down deps

# Default target
help:
	@echo "Financial Transaction System - Available Commands:"
	@echo "=================================================="
	@echo "  build         - Build the application"
	@echo "  run           - Run the application locally"
	@echo "  test          - Run tests"
	@echo "  test-api      - Test API endpoints"
	@echo "  clean         - Clean build artifacts"
	@echo "  docker-up     - Start Docker services"
	@echo "  docker-down   - Stop Docker services"
	@echo "  migrate-up    - Run database migrations"
	@echo "  migrate-down  - Rollback database migrations"
	@echo "  deps          - Download dependencies"

# Build the application
build:
	@echo "Building application..."
	go build -o bin/server cmd/server/main.go
	go build -o bin/migrate cmd/migrate/main.go

# Run the application
run:
	@echo "Starting application..."
	go run cmd/server/main.go

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Test API endpoints
test-api:
	@echo "Testing API endpoints..."
	./test_api.sh

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	go clean

# Start Docker services
docker-up:
	@echo "Starting Docker services..."
	docker-compose up -d
	@echo "Waiting for services to be ready..."
	sleep 10

# Stop Docker services
docker-down:
	@echo "Stopping Docker services..."
	docker-compose down

# Run database migrations
migrate-up:
	@echo "Running database migrations..."
	go run cmd/migrate/main.go up

# Rollback database migrations
migrate-down:
	@echo "Rolling back database migrations..."
	go run cmd/migrate/main.go down

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

# Development workflow
dev: docker-up migrate-up run

# Full setup
setup: deps docker-up migrate-up
	@echo "Setup completed! You can now run 'make run' to start the server."

# Docker build
docker-build:
	@echo "Building Docker image..."
	docker build -t financial-transaction-system .

# Run with Docker
docker-run: docker-build
	@echo "Running with Docker..."
	docker-compose up --build 