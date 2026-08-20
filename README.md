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

## 🏛️ System Architecture

```mermaid
flowchart TD
    subgraph MCP_Host["💻 MCP Host (AI Client / IDE)"]
        IDE["Harness or AI-IDE\n(Claude Desktop, Cursor, VS Code, Opencode)"]
        Agent["AI Agent Model\n(Claude 3.5 Sonnet, Gemini 1.5 Pro, GPT-4o)"]
        IDE --> Agent
    end

    MCP_Host -- "JSON-RPC 2.0 (stdio / SSE)" --> Server_Boundary

    subgraph Server_Boundary["⚙️ meshery-mcp-server (Golang Engine)"]
        direction TB
        
        subgraph Transport_Layer["Transport Layer"]
            StdioTransport["stdio Pipe Transport"]
            SSETransport["Streamable SSE HTTP Transport"]
        end

        subgraph Core_Engine["Core Engine & Registrant Seam"]
            Registrant["Registration Seam\n(Tools | Resources | Prompts)"]
            Tools["Read-Only Tools\n(list_designs, server_info)"]
            Resources["MeshSync Resources\n(meshsync://topology)"]
            Prompts["Design Prompts\n(pattern_analysis)"]
            Registrant --> Tools
            Registrant --> Resources
            Registrant --> Prompts
        end

        subgraph Security_Error_Layer["Security & Error Boundary Layer"]
            Sanitizer["🔒 Response Boundary Sanitizer (pkg/security)\n- Immutability / Non-mutation\n- Nested JSON/Map Redaction\n- Precision Key Matching\n- Error-Path Scrubbing"]
            MeshKitErrors["⚠️ MeshKit Error Engine (pkg/errors)\n- Structured Error Codes (1001-MCP, 1002-MCP)\n- Non-leaking Remediation Guidance"]
        end

        subgraph Client_Boundary["Meshery Client Boundary"]
            SharedClient["Decoupled MesheryClient Interface\n(pkg/meshery/client.go)"]
        end

        Transport_Layer --> Core_Engine
        Core_Engine --> Sanitizer
        Sanitizer --> MeshKitErrors
        MeshKitErrors --> SharedClient
    end

    SharedClient -- "Authenticated REST / GraphQL API" --> Meshery_Core

    subgraph Meshery_Core["☁️ CNCF Meshery Core Infrastructure"]
        MesheryAPI["Meshery Server REST API\n(Port 9081)"]
        MeshSync["MeshSync Discovery Engine"]
        Kubernetes["Kubernetes & Cloud Infrastructure"]
        
        MesheryAPI --> MeshSync
        MeshSync --> Kubernetes
    end

    style Sanitizer fill:#1b4332,stroke:#2d6a4f,stroke-width:2px,color:#fff
    style MeshKitErrors fill:#3d0066,stroke:#5c0099,stroke-width:2px,color:#fff
    style Server_Boundary fill:#111827,stroke:#374151,stroke-width:2px,color:#fff
```

---

## 🚀 Overview

`meshery-mcp-server` is a Golang implementation of the **Model Context Protocol (MCP)** for [CNCF Meshery](https://meshery.io). It acts as a secure bridge connecting AI-assisted development tools (**Claude Desktop**, **Cursor IDE**, **VS Code AI agents**) directly to the Meshery infrastructure engine and MeshSync Kubernetes topology discovery.

---

## ✨ Key Technical Capabilities

* **Dual Transport Architecture**: Supports local `stdio` (for IDE process pipes) and streamable `SSE` (Server-Sent Events over HTTP).
* **Response Boundary Secret Sanitizer (`pkg/security`)**: SDK-agnostic recursive JSON redactor scrubbing tokens, passwords, and KubeConfig secrets before streaming context to AI models.
* **MeshKit Error Code Framework (`pkg/errors`)**: Implements structured error codes (`1001-MCP`, `1002-MCP`) compliant with Layer5 MeshKit standards.
* **Sub-Microsecond Benchmarks (`pkg/security/sanitizer_bench_test.go`)**: High-performance execution running at **1,363,168 ops/sec** (<883 ns/op).
* **Decoupled Architecture**: `MesheryClient` narrow interface pattern keeping tool handlers lightweight and transport-agnostic.
* **100% Test-Driven**: 8 comprehensive unit test cases covering immutability, error-path scrubbing, precision key matching, and nil value resilience.
* **Interactive Walkthrough**: Complete demo and execution guide in [`DEMO.md`](DEMO.md).

---

## 🏎️ Performance Benchmarks

```text
goos: linux
goarch: amd64
pkg: github.com/peaush07/meshery-mcp-server-poc/pkg/security
cpu: 13th Gen Intel(R) Core(TM) i5-13420H
BenchmarkSanitizeMap-12     	 1363168	       882.4 ns/op
BenchmarkSanitizeJSON-12    	  319851	      3449 ns/op
PASS
ok  	github.com/peaush07/meshery-mcp-server-poc/pkg/security	3.239s
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

## 🛠️ Getting Started & Building

```bash
# Clone the repository
git clone https://github.com/peaush07/meshery-mcp-server-poc.git
cd meshery-mcp-server-poc

# Run Unit Test Suite
make test

# Run Performance Benchmarks
make bench

# Build Binary
make build

# Build Docker Container Image
make docker-build

# Run in Stdio mode
./bin/meshery-mcp-server -transport=stdio

# Run in SSE HTTP mode
./bin/meshery-mcp-server -transport=sse -port=8080
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
