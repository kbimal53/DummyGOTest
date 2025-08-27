.PHONY: run build test clean help migrate

# Default target
help:
	@echo "Available commands:"
	@echo "  make run       - Run the application"
	@echo "  make build     - Build the application"
	@echo "  make test      - Run tests"
	@echo "  make migrate   - Run database migration"
	@echo "  make clean     - Clean build artifacts"
	@echo "  make help      - Show this help message"

# Run the application
run:
	@echo "🚀 Starting the Go API server..."
	go run main.go database.go

# Build the application
build:
	@echo "🔨 Building the application..."
	go build -o api main.go database.go
	@echo "✅ Build complete! Binary: ./api"

# Run database migration
migrate:
	@echo "🗄️  Running database migration..."
	go run migrate.go database.go

# Run tests
test:
	@echo "🧪 Running tests..."
	go test -v

# Run tests with coverage
test-coverage:
	@echo "🧪 Running tests with coverage..."
	go test -v -cover

# Install dependencies
deps:
	@echo "📦 Installing dependencies..."
	go mod tidy

# Clean build artifacts
clean:
	@echo "🧹 Cleaning build artifacts..."
	rm -f api
	go clean

# Format code
fmt:
	@echo "🎨 Formatting code..."
	go fmt ./...

# Run linter (requires golangci-lint)
lint:
	@echo "🔍 Running linter..."
	golangci-lint run

# Development mode with auto-reload (requires air)
dev:
	@echo "🔄 Starting development mode with auto-reload..."
	air
