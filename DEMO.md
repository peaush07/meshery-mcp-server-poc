# 🚀 Meshery MCP Server PoC — Interactive Walkthrough & Demo Guide

> **Project:** CNCF Meshery — Model Context Protocol (MCP) Server  
> **Applicant:** Peaush Paul ([@peaush07](https://github.com/peaush07)) | **LFX Mentorship 2026 Term 3**  
> **Issue Reference:** [meshery/meshery#19446](https://github.com/meshery/meshery/issues/19446)

---

## 📌 Demonstration Overview

This document provides a complete technical walkthrough demonstrating the **`meshery-mcp-server`** Proof of Concept (PoC) in action across dual transports (`stdio` and `SSE`), featuring automatic secret redaction and JSON-RPC 2.0 execution.

---

## 🛠️ 1. Building the Binary & Running Unit Tests

### Build Binary
```bash
export PATH=$PATH:/home/peaush/Downloads/go/bin
cd ~/meshery-mcp-server-poc
make build
```

### Run Unit Test Suite
Verify that the `pkg/security` secret sanitizer handles nested maps, slices, and case variations with **100% PASS**:

```bash
make test
```

**Expected Test Output:**
```text
=== RUN   TestSanitizeMap_SensitiveKeys
--- PASS: TestSanitizeMap_SensitiveKeys (0.00s)
=== RUN   TestSanitizeMap_NestedStructures
--- PASS: TestSanitizeMap_NestedStructures (0.00s)
=== RUN   TestSanitizeMap_CaseInsensitive
--- PASS: TestSanitizeMap_CaseInsensitive (0.00s)
=== RUN   TestSanitizeJSON_ValidJSON
--- PASS: TestSanitizeJSON_ValidJSON (0.00s)
PASS
ok  	github.com/peaush07/meshery-mcp-server-poc/pkg/security	0.003s
```

---

## 💻 2. Testing `stdio` Transport (Claude Desktop / Cursor IDE)

In `stdio` mode, the server opens standard I/O pipes. All internal debug logs are piped exclusively to `os.Stderr` to maintain a zero-corruption JSON-RPC 2.0 channel on `os.Stdout`.

### Start Server in Stdio Mode
```bash
./bin/meshery-mcp-server -transport=stdio
```

### Step 1: Protocol Handshake (`initialize`)
Send JSON payload to `stdin`:
```json
{"jsonrpc":"2.0","id":1,"method":"initialize"}
```

**Server Response on `stdout`:**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "capabilities": {
      "prompts": {},
      "resources": {},
      "tools": {}
    },
    "protocolVersion": "2024-11-05",
    "serverInfo": {
      "name": "meshery-mcp-server",
      "version": "v0.1.0-poc"
    }
  }
}
```

### Step 2: List Available Tools (`tools/list`)
```json
{"jsonrpc":"2.0","id":2,"method":"tools/list"}
```

**Server Response on `stdout`:**
```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "result": {
    "tools": [
      {
        "name": "list_designs",
        "description": "Lists infrastructure design patterns stored on the connected Meshery server."
      },
      {
        "name": "run_nighthawk_test",
        "description": "Executes a Nighthawk performance load test benchmark against target infrastructure."
      }
    ]
  }
}
```

### Step 3: Execute Tool with Automatic Secret Redaction (`tools/call`)
```json
{
  "jsonrpc": "2.0",
  "id": 3,
  "method": "tools/call",
  "params": {
    "name": "list_designs"
  }
}
```

**Sanitized Response Output:**
```json
{
  "jsonrpc": "2.0",
  "id": 3,
  "result": {
    "content": [
      {
        "type": "text",
        "text": "Retrieved 1 Meshery design pattern(s).\nCredentials Redacted: [REDACTED_SECRET]"
      }
    ]
  }
}
```

---

## 🌐 3. Testing `SSE` HTTP Mode (Remote / Web Server)

Start HTTP server listening on port `8080`:

```bash
./bin/meshery-mcp-server -transport=sse -port=8080 -meshery-url=http://localhost:9081
```

**Server Log on `stderr`:**
```text
2026/08/19 00:50:00 [Meshery-MCP-Server] Initializing server... (transport: sse)
2026/08/19 00:50:00 [Meshery-MCP-Server] Starting SSE HTTP server on :8080...
```

* **Server-Sent Events Stream:** `GET http://localhost:8080/sse`
* **JSON-RPC HTTP Endpoint:** `POST http://localhost:8080/message`

---

## 🔒 4. Security Architecture (`pkg/security/sanitizer.go`)

To prevent accidental leakage of Kubernetes API tokens, service account credentials, or database passwords to AI context windows, the PoC implements an interceptor pattern at the client response boundary:

```
[ Raw API Payload ] --> [ SanitizeMap() Interceptor ] --> [ Masked Payload ] --> [ MCP Tool Output ]
                             |
                             +-- Checks: token, password, secret, kubeconfig, private_key
```

* Recursively traverses nested JSON maps (`map[string]interface{}`) and slices (`[]interface{}`).
* Performs case-insensitive key matching (`AuthToken`, `PASSWORD`, `KubeConfig`).
* Replaces matching secret values with `"[REDACTED_SECRET]"`.

---

## 🤝 5. IDE Client Configuration (Cursor / Claude Desktop)

Add the following configuration block to `claude_desktop_config.json` or Cursor MCP settings:

```json
{
  "mcpServers": {
    "meshery": {
      "command": "/home/peaush/meshery-mcp-server-poc/bin/meshery-mcp-server",
      "args": [
        "-transport=stdio",
        "-meshery-url=http://localhost:9081"
      ]
    }
  }
}
```
