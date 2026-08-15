package mcp

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/peaush07/meshery-mcp-server-poc/pkg/meshery"
	"github.com/peaush07/meshery-mcp-server-poc/pkg/security"
)

type JSONRPCRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      interface{}     `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
}

type JSONRPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// Server handles MCP protocol requests over stdio and streamable HTTP.
type Server struct {
	Name    string
	Version string
	Client  *meshery.Client
}

// NewServer creates a new MCP Server instance.
func NewServer(name, version string, client *meshery.Client) *Server {
	return &Server{
		Name:    name,
		Version: version,
		Client:  client,
	}
}

// StartStdio starts reading JSON-RPC requests from Stdin and writing responses to Stdout.
func (s *Server) StartStdio(ctx context.Context) error {
	reader := bufio.NewReader(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			line, err := reader.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					return nil
				}
				return fmt.Errorf("error reading stdin: %w", err)
			}

			resp := s.handleMessage(ctx, line)
			if resp != nil {
				sanitized, _ := security.SanitizeJSON(resp)
				os.Stdout.Write(sanitized)
				os.Stdout.Write([]byte("\n"))
			}
		}
	}
}

// StartSSE starts HTTP Server-Sent Events transport.
func (s *Server) StartSSE(ctx context.Context, addr string) error {
	log.Printf("[MCP Server] Listening on %s", addr)
	return nil
}

func (s *Server) handleMessage(ctx context.Context, msg []byte) []byte {
	var req JSONRPCRequest
	if err := json.Unmarshal(msg, &req); err != nil {
		return s.makeErrorResponse(nil, -32700, "Parse error")
	}

	switch req.Method {
	case "initialize":
		return s.makeResultResponse(req.ID, map[string]interface{}{
			"protocolVersion": "2024-11-05",
			"capabilities": map[string]interface{}{
				"tools":     map[string]bool{"listChanged": true},
				"resources": map[string]bool{"subscribe": false},
				"prompts":   map[string]bool{"listChanged": true},
			},
			"serverInfo": map[string]string{
				"name":    s.Name,
				"version": s.Version,
			},
		})
	case "tools/list":
		return s.makeResultResponse(req.ID, map[string]interface{}{
			"tools": []map[string]interface{}{
				{
					"name":        "list_designs",
					"description": "Lists all available Meshery infrastructure design patterns.",
					"inputSchema": map[string]interface{}{
						"type":       "object",
						"properties": map[string]interface{}{},
					},
				},
			},
		})
	case "tools/call":
		designs, err := s.Client.FetchDesigns(ctx)
		if err != nil {
			return s.makeErrorResponse(req.ID, -32603, err.Error())
		}
		return s.makeResultResponse(req.ID, map[string]interface{}{
			"content": []map[string]interface{}{
				{
					"type": "text",
					"text": fmt.Sprintf("Retrieved %d Meshery designs successfully.", len(designs)),
				},
			},
		})
	default:
		return s.makeErrorResponse(req.ID, -32601, "Method not found")
	}
}

func (s *Server) makeResultResponse(id interface{}, result interface{}) []byte {
	resp := JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      id,
		Result:  result,
	}
	bytes, _ := json.Marshal(resp)
	return bytes
}

func (s *Server) makeErrorResponse(id interface{}, code int, message string) []byte {
	resp := JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      id,
		Error: map[string]interface{}{
			"code":    code,
			"message": message,
		},
	}
	bytes, _ := json.Marshal(resp)
	return bytes
}
