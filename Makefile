# Makefile for go-for-rays - Ray Tracer Challenge in Go

# Variables
BINARY_NAME=go-for-rays
MAIN_PATH=./main.go

# Default target
.PHONY: all
all: clean test build

# Build the binary
.PHONY: build
build:
	@echo "Building $(BINARY_NAME)..."
	go build -o bin/$(BINARY_NAME) $(MAIN_PATH)

# Build for multiple platforms
.PHONY: build-all
build-all: build-linux build-windows build-darwin

.PHONY: build-linux
build-linux:
	@echo "Building for Linux..."
	GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY_NAME)-linux-amd64 $(MAIN_PATH)

.PHONY: build-windows
build-windows:
	@echo "Building for Windows..."
	GOOS=windows GOARCH=amd64 go build -o bin/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PATH)

.PHONY: build-darwin
build-darwin:
	@echo "Building for macOS..."
	GOOS=darwin GOARCH=amd64 go build -o bin/$(BINARY_NAME)-darwin-amd64 $(MAIN_PATH)

# Run the application
.PHONY: run
run:
	@echo "Running $(BINARY_NAME)..."
	go run $(MAIN_PATH)

# Run with arguments (use: make run-args ARGS="--help")
.PHONY: run-args
run-args:
	@echo "Running $(BINARY_NAME) with args: $(ARGS)"
	go run $(MAIN_PATH) $(ARGS)

# Download dependencies
.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run tests with race detection
.PHONY: test-race
test-race:
	@echo "Running tests with race detection..."
	go test -race -v ./...

# Benchmark tests
.PHONY: bench
bench:
	@echo "Running benchmarks..."
	go test -bench=. -benchmem ./...

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Lint code
.PHONY: lint
lint:
	@echo "Linting code..."
	golangci-lint run

# Vet code
.PHONY: vet
vet:
	@echo "Vetting code..."
	go vet ./...

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -f coverage.out coverage.html

# Run a specific test package
.PHONY: test-pkg
test-pkg:
	@echo "Running tests for package: $(PKG)"
	go test -v ./$(PKG)

# Install the binary globally
.PHONY: install
install:
	@echo "Installing $(BINARY_NAME) globally..."
	go install $(MAIN_PATH)

# Help target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  all              - Clean, test, and build"
	@echo "  build            - Build the binary"
	@echo "  build-all        - Build for multiple platforms"
	@echo "  run              - Run the application"
	@echo "  run-args         - Run with arguments (use ARGS='...')"
	@echo "  deps             - Download dependencies"
	@echo "  test             - Run tests"
	@echo "  test-coverage    - Run tests with coverage report"
	@echo "  test-race        - Run tests with race detection"
	@echo "  bench            - Run benchmarks"
	@echo "  fmt              - Format code"
	@echo "  lint             - Lint code"
	@echo "  vet              - Vet code"
	@echo "  clean            - Clean build artifacts"
	@echo "  test-pkg         - Run tests for specific package (use PKG='...')"
	@echo "  install          - Install binary globally"
	@echo "  help             - Show this help message"
