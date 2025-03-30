# Default target: Run the application
run:
	@echo "Running the application..."
	go run cmd/main.go

# Run the application inside Docker
dev:
	@echo "Running the application in Docker..."
	docker-compose up -d && go run cmd/main.go

# Stop all Docker containers
down:
	@echo "Stopping all containers..."
	docker-compose down

# Generate API documentation with swag
swag:
	@echo "Generating API documentation..."
	swag init -g cmd/main.go -o docs

test:
	@echo "Running tests..."
	go test ./...

test-cover:
	@echo "Running tests with coverage..."
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

# Show help message
help:
	@echo "Makefile for Go project"
	@echo "Available targets:"
	@echo "  run        - Run the application"
	@echo "  dev        - Run the application in Docker"
	@echo "  down       - Stop all Docker containers"
	@echo "  swag       - Generate API documentation using swag"
	@echo "  test       - Run all tests"
	@echo "  test-cover - Run all tests with coverage"
	@echo "  help       - Show this message"
