# Meshery MCP Server (Proof of Concept)

[![LFX Mentorship](https://img.shields.io/badge/LFX-Mentorship%202026%20Term%203-blue.svg)](https://mentorship.lfx.linuxfoundation.org/)
[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8.svg)](https://golang.org/)

**Author:** Peaush Paul ([@peaush07](https://github.com/peaush07))  
**Email:** peaushpaul99@gmail.com  
**CNCF Slack:** `@PeaushPaul`  

This repository contains a **Proof of Concept (PoC) implementation of the Meshery Model Context Protocol (MCP) Server**, created as part of the LFX Mentorship application for CNCF Meshery (2026 Term 3, Issue #19446).

## Overview

`meshery-mcp-server` bridges CNCF Meshery's cloud-native infrastructure management plane to AI coding assistants (Claude Desktop, Cursor IDE) via the Model Context Protocol (MCP).

### Features Implemented in PoC

- **Golang Architecture**: Modular architecture (`pkg/mcp`, `pkg/meshery`, `pkg/security`).
- **MCP Protocol Engine**: Handles JSON-RPC 2.0 initialization, tools list, and execution.
- **Stdio & SSE Transports**: Supports both standard I/O (for local IDEs) and HTTP Server-Sent Events.
- **Security Sanitizer**: Automatic masking of sensitive tokens, passwords, and K8s secrets.
- **Meshery Client Wrapper**: Connects to Meshery REST API endpoints (`/api/content/patterns`).

## Getting Started

### Prerequisites

- Go 1.22+
- Make

### Building and Running

```bash
# Clone the repository
git clone https://github.com/peaush07/meshery-mcp-server-poc.git
cd meshery-mcp-server-poc

# Build the binary
make build

# Run in Stdio mode
./bin/meshery-mcp-server -transport=stdio

# Run in SSE HTTP mode
./bin/meshery-mcp-server -transport=sse -port=8080
```

## Testing with Claude Desktop / Cursor

Add the following to your `claude_desktop_config.json` or Cursor MCP settings:

```json
{
  "mcpServers": {
    "meshery": {
      "command": "/path/to/meshery-mcp-server-poc/bin/meshery-mcp-server",
      "args": ["-transport=stdio", "-meshery-url=http://localhost:9081"]
    }
  }
}
```

---
*Created by Peaush Paul for the CNCF Meshery LFX Mentorship 2026 Term 3.*
