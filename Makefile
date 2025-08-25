# Mango Password Manager Makefile

# Variables
BINARY_NAME=mango
MAIN_FILE=main.go
BUILD_DIR=build
VERSION=$(shell git describe --tags --always --dirty)
LDFLAGS=-ldflags="-s -w -X main.version=$(VERSION)"

# Default target
.PHONY: all
all: build

# Build for current platform
.PHONY: build
build:
	@echo "Building $(BINARY_NAME)..."
	go build $(LDFLAGS) -o $(BINARY_NAME) $(MAIN_FILE)
	@echo "✅ Build complete: $(BINARY_NAME)"

# Build optimized version
.PHONY: build-optimized
build-optimized:
	@echo "Building optimized $(BINARY_NAME)..."
	go build -ldflags="-s -w" -o $(BINARY_NAME) $(MAIN_FILE)
	@echo "✅ Optimized build complete: $(BINARY_NAME)"

# Cross-compile for all platforms
.PHONY: build-all
build-all: clean
	@echo "Cross-compiling for all platforms..."
	@mkdir -p $(BUILD_DIR)
	
	# Linux AMD64
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_FILE)
	
	# Linux ARM64
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 $(MAIN_FILE)
	
	# macOS AMD64
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_FILE)
	
	# macOS ARM64 (Apple Silicon)
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN_FILE)
	
	# Windows AMD64
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_FILE)
	
	# Windows ARM64
	GOOS=windows GOARCH=arm64 go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-arm64.exe $(MAIN_FILE)
	
	@echo "✅ Cross-compilation complete! Binaries in $(BUILD_DIR)/"
	@ls -la $(BUILD_DIR)/

# Install to system PATH
.PHONY: install
install: build
	@echo "Installing $(BINARY_NAME) to /usr/local/bin..."
	sudo mv $(BINARY_NAME) /usr/local/bin/
	@echo "✅ $(BINARY_NAME) installed successfully!"

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	go test -v ./...

# Run with race detection
.PHONY: test-race
test-race:
	@echo "Running tests with race detection..."
	go test -race -v ./...

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	go fmt ./...
	@echo "✅ Code formatted"

# Lint code
.PHONY: lint
lint:
	@echo "Linting code..."
	golangci-lint run
	@echo "✅ Linting complete"

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_NAME)-*
	rm -rf $(BUILD_DIR)
	@echo "✅ Clean complete"

# Show help
.PHONY: help
help:
	@echo "Mango Password Manager - Build Commands"
	@echo ""
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build          Build for current platform"
	@echo "  build-optimized Build optimized version (smaller size)"
	@echo "  build-all      Cross-compile for all platforms"
	@echo "  install        Install to system PATH (/usr/local/bin)"
	@echo "  test           Run tests"
	@echo "  test-race      Run tests with race detection"
	@echo "  fmt            Format code"
	@echo "  lint           Lint code (requires golangci-lint)"
	@echo "  clean          Clean build artifacts"
	@echo "  help           Show this help message"
	@echo ""
	@echo "Examples:"
	@echo "  make build              # Build for current platform"
	@echo "  make build-all          # Build for all platforms"
	@echo "  make install            # Build and install to PATH"

# Development helpers
.PHONY: dev
dev:
	@echo "Running in development mode..."
	go run $(MAIN_FILE) --help

.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy
	@echo "✅ Dependencies updated"