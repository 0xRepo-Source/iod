# Makefile for IndexOf Downloader

.PHONY: build clean test install run help

# Build the application
build:
	go build -o iod.exe .

# Build for multiple platforms
build-all:
	GOOS=windows GOARCH=amd64 go build -o bin/iod-windows-amd64.exe .
	GOOS=linux GOARCH=amd64 go build -o bin/iod-linux-amd64 .
	GOOS=darwin GOARCH=amd64 go build -o bin/iod-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -o bin/iod-darwin-arm64 .

# Clean build artifacts
clean:
	del /Q iod.exe 2>nul || exit 0
	del /Q /S bin\* 2>nul || exit 0

# Run tests
test:
	go test ./... -v

# Install dependencies
deps:
	go mod download
	go mod tidy

# Install the binary globally
install:
	go install .

# Run the application (example)
run:
	go run . list https://example.com/pdf/microsoft/ -d 0

# Display help
help:
	@echo Available targets:
	@echo   build      - Build the application for current platform
	@echo   build-all  - Build for all platforms (Windows, Linux, macOS)
	@echo   clean      - Remove build artifacts
	@echo   test       - Run tests
	@echo   deps       - Download and tidy dependencies
	@echo   install    - Install the binary globally
	@echo   run        - Run example command
	@echo   help       - Show this help message

