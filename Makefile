.PHONY: build test bench docker-build run-stdio run-sse clean

BINARY_NAME=bin/meshery-mcp-server

build:
	@echo "Building $(BINARY_NAME) binary..."
	@mkdir -p bin
	@go build -o $(BINARY_NAME) main.go

test:
	@echo "Running Go unit test suite..."
	@go test -v ./...

bench:
	@echo "Running Go benchmark suite..."
	@go test -bench=. ./pkg/security

docker-build:
	@echo "Building Docker container image..."
	@docker build -t meshery-mcp-server:latest .

run-stdio: build
	@echo "Starting meshery-mcp-server in stdio transport mode..."
	@./$(BINARY_NAME) -transport=stdio

run-sse: build
	@echo "Starting meshery-mcp-server in SSE HTTP transport mode..."
	@./$(BINARY_NAME) -transport=sse -port=8080

clean:
	@rm -rf bin/
