.PHONY: build run test lint clean

BINARY_NAME=meshery-mcp-server

build:
	@echo "Building meshery-mcp-server binary..."
	go build -o bin/$(BINARY_NAME) main.go

run: build
	@echo "Running meshery-mcp-server..."
	./bin/$(BINARY_NAME)

test:
	@echo "Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...

lint:
	@echo "Running golangci-lint..."
	golangci-lint run

clean:
	rm -rf bin/ coverage.out
