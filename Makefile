.PHONY: run test build

# Load environment variables from .env file
include .env
export

# Run the application
run:
	@echo "Starting application..."
	@go run cmd/api/main.go

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Build the application
build:
	@echo "Building application..."
	@go build -o bin/api cmd/api/main.go
