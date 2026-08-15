package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/peaush07/meshery-mcp-server-poc/pkg/mcp"
	"github.com/peaush07/meshery-mcp-server-poc/pkg/meshery"
)

func main() {
	var (
		mesheryURL = flag.String("meshery-url", "http://localhost:9081", "Meshery Server Endpoint")
		transport  = flag.String("transport", "stdio", "Transport mode: stdio or sse")
		port       = flag.Int("port", 8080, "Port for HTTP SSE transport mode")
	)
	flag.Parse()

	// Redirect normal log output to Stderr so Stdio JSON-RPC protocol is clean
	log.SetOutput(os.Stderr)
	log.Printf("[Meshery-MCP-Server] Initializing server... (transport: %s)", *transport)

	// Create Meshery API Client wrapper
	client := meshery.NewClient(*mesheryURL)

	// Create MCP Server instance
	server := mcp.NewServer("meshery-mcp-server", "v0.1.0-poc", client)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	switch *transport {
	case "stdio":
		log.Println("[Meshery-MCP-Server] Starting stdio transport...")
		if err := server.StartStdio(ctx); err != nil {
			log.Fatalf("Stdio server error: %v", err)
		}
	case "sse":
		log.Printf("[Meshery-MCP-Server] Starting SSE HTTP server on port %d...", *port)
		if err := server.StartSSE(ctx, fmt.Sprintf(":%d", *port)); err != nil {
			log.Fatalf("SSE server error: %v", err)
		}
	default:
		log.Fatalf("Unsupported transport: %s", *transport)
	}
}
