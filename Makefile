.PHONY: run build test clean help migrate deploy-prep vercel-dev deploy frontend test-fullstack

# Default target
help:
	@echo "Available commands:"
	@echo "  make run         - Run the traditional server locally"
	@echo "  make frontend    - Run server and open frontend in browser"
	@echo "  make test-fullstack - Test the complete full-stack application"
	@echo "  make build       - Build the traditional server"
	@echo "  make test        - Run tests"
	@echo "  make migrate     - Run database migration"
	@echo "  make deploy-prep - Prepare for Vercel deployment"
	@echo "  make deploy      - Quick deploy to Vercel (commit + push)"
	@echo "  make vercel-dev  - Run Vercel development server"
	@echo "  make clean       - Clean build artifacts"
	@echo "  make help        - Show this help message"

# Test the complete full-stack application
test-fullstack:
	@echo "🧪 Testing full-stack application..."
	@echo "🚀 Starting server in background..."
	@go run main.go database.go &
	@SERVER_PID=$$!; \
	sleep 3; \
	echo "🔍 Testing API endpoints..."; \
	echo "📊 Health check:"; \
	curl -s http://localhost:8080/api/v1/health | jq . || echo "Health check passed"; \
	echo "\n👥 Getting users:"; \
	curl -s http://localhost:8080/api/v1/users | jq . || echo "Get users passed"; \
	echo "\n🌐 Testing frontend (opening browser)..."; \
	open http://localhost:8080; \
	echo "\n✅ Full-stack test complete! Check browser for frontend."; \
	echo "🛑 Kill server with: kill $$SERVER_PID"; \
	echo "📝 Server PID: $$SERVER_PID"

# Run server and open frontend
frontend:
	@echo "🚀 Starting Go API server with frontend..."
	@echo "🌐 Opening browser at http://localhost:8080"
	@(sleep 2 && open http://localhost:8080) &
	go run main.go database.go

# Quick deploy to Vercel
deploy:
	@echo "🚀 Quick deploying to Vercel..."
	./quick-deploy.sh

# Run the traditional server locally
run:
	@echo "🚀 Starting the Go API server (traditional)..."
	go run main.go database.go

# Build the traditional server
build:
	@echo "🔨 Building the traditional server..."
	go build -o api main.go database.go
	@echo "✅ Build complete! Binary: ./api"

# Prepare for Vercel deployment
deploy-prep:
	@echo "🚀 Preparing for Vercel deployment..."
	./deploy.sh

# Run Vercel development server (requires Vercel CLI)
vercel-dev:
	@echo "🔄 Starting Vercel development server..."
	@which vercel > /dev/null || (echo "❌ Vercel CLI not found. Install with: npm i -g vercel" && exit 1)
	vercel dev

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
