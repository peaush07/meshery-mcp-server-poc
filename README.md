# 🛠️ Meshery MCP Server Proof of Concept (PoC)

[![CI Build & Test](https://github.com/peaush07/meshery-mcp-server-poc/actions/workflows/ci.yml/badge.svg)](https://github.com/peaush07/meshery-mcp-server-poc/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/badge/Go-1.22%2B-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](LICENSE)
[![Mentorship](https://img.shields.io/badge/LFX_Mentorship-2026_Term_3-green.svg)](https://mentorship.lfx.linuxfoundation.org/)

> **Official LFX Mentorship Proposal Reference Implementation**  
> **Project:** CNCF - Meshery: MCP Server  
> **Issue:** [meshery/meshery#19446](https://github.com/meshery/meshery/issues/19446)  
> **Applicant:** Peaush Paul ([@peaush07](https://github.com/peaush07)) | Kolkata, India (UTC+5:30)

---

## 🚀 Overview

`meshery-mcp-server` is a Golang implementation of the **Model Context Protocol (MCP)** for [CNCF Meshery](https://meshery.io). It acts as a secure bridge connecting AI-assisted development tools (**Claude Desktop**, **Cursor IDE**, **VS Code AI agents**) directly to the Meshery infrastructure engine and MeshSync Kubernetes topology discovery.

---

## ✨ Key Technical Capabilities

* **Dual Transport Architecture**: Supports local `stdio` (for IDE process pipes) and streamable `SSE` (Server-Sent Events over HTTP).
* **Response Boundary Secret Sanitizer (`pkg/security`)**: SDK-agnostic recursive JSON redactor scrubbing tokens, passwords, and KubeConfig secrets before streaming context to AI models.
* **Decoupled Architecture**: `MesheryClient` narrow interface pattern keeping tool handlers lightweight and agnostic of transport/auth details.
* **100% Test-Driven**: 8 comprehensive unit test cases covering immutability/non-mutation, error-path scrubbing, precision key matching, and nil value resilience.
* **Interactive Walkthrough**: Complete demo and execution guide in [`DEMO.md`](DEMO.md).

---

## 🛠️ Getting Started & Prerequisites

### Prerequisites
* **Go 1.22+**
* **Make**

### Building and Running

```bash
# 1. Clone the repository
git clone https://github.com/peaush07/meshery-mcp-server-poc.git
cd meshery-mcp-server-poc

# 2. Run Unit Test Suite (100% PASS)
make test

# 3. Build Binary
make build

# 4. Run in Stdio mode (for Cursor / Claude Desktop)
./bin/meshery-mcp-server -transport=stdio

# 5. Run in SSE HTTP mode (for web/remote streaming on port 8080)
./bin/meshery-mcp-server -transport=sse -port=8080
```

---

## 🧪 Unit Test Suite Verification (`pkg/security/sanitizer_test.go`)

```text
=== RUN   TestSanitizeMap_SensitiveKeys            --- PASS (0.00s)
=== RUN   TestSanitizeMap_NestedStructures          --- PASS (0.00s)
=== RUN   TestSanitizeMap_CaseInsensitive           --- PASS (0.00s)
=== RUN   TestSanitizeJSON_ValidJSON               --- PASS (0.00s)
=== RUN   TestSanitizeMap_Immutability               --- PASS (0.00s)
=== RUN   TestSanitizeMap_PrecisionKeyMatching      --- PASS (0.00s)
=== RUN   TestSanitizeString_ErrorPathRedaction     --- PASS (0.00s)
=== RUN   TestSanitizeMap_NilOrEmptyHandling        --- PASS (0.00s)
PASS
ok  	github.com/peaush07/meshery-mcp-server-poc/pkg/security	0.002s
```

---

## 💻 Claude Desktop / Cursor IDE Configuration

Add the following to your `claude_desktop_config.json` or Cursor MCP settings:

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

---

## 📄 License & Community

This project is licensed under the Apache License 2.0. Built for the CNCF Meshery & Layer5 open source community.
