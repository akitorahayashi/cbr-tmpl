# ==============================================================================
# justfile for cbr-tmpl automation
# ==============================================================================

# default target
default: help

# Show available recipes
help:
    @just --list

# ==============================================================================
# Environment Setup
# ==============================================================================

# Initialize project: install dependencies
setup:
    @echo "📦 Downloading Go dependencies..."
    @go mod download

# ==============================================================================
# Development Commands
# ==============================================================================

# Run the CLI application
run *args:
    @go run ./cmd/cbr-tmpl {{args}}

# Build the CLI binary
build:
    @echo "🔨 Building cbr-tmpl..."
    @go build -o bin/cbr-tmpl ./cmd/cbr-tmpl
    @echo "✅ Built to bin/cbr-tmpl"

# ==============================================================================
# CODE QUALITY
# ==============================================================================

# Automatically format and fix code
fix:
    @echo "🔧 Formatting and fixing code..."
    @goimports -w .
    @golangci-lint run --fix

# Run static checks (format check, lint, vet)
check:
    @echo "🔍 Running static checks..."
    @test -z "$$(goimports -l .)" || (echo "Run 'just fix' to format code" && exit 1)
    @golangci-lint run
    @go vet ./...

# ==============================================================================
# TESTING
# ==============================================================================

# Run all tests
test:
    @echo "🚀 Running tests..."
    @go test ./...
    @echo "✅ All tests passed!"

# ==============================================================================
# CLEANUP
# ==============================================================================

# Remove build artifacts
clean:
    @echo "🧹 Cleaning up..."
    @rm -rf bin/
    @go clean -cache
    @echo "✅ Done"
