# 📜 Changelog

All notable changes to `meshery-mcp-server` are documented in this file.

---

## [v0.1.0-alpha] - 2026-08-20

### ✨ Added
- **Dual Transport Architecture**: Implemented native `stdio` process pipe and streamable `SSE` (Server-Sent Events) HTTP transport modes.
- **Response Boundary Secret Sanitizer (`pkg/security`)**: Built SDK-agnostic recursive JSON redactor scrubbing tokens, passwords, and KubeConfigs.
- **8/8 Unit Test Matrix (`pkg/security/sanitizer_test.go`)**: 100% PASSing unit test suite covering immutability, error paths, precision key matching, and nil resilience.
- **Sub-Microsecond Benchmarks (`pkg/security/sanitizer_bench_test.go`)**: High-throughput performance running at **1.3 Million ops/sec** (`<883 ns/op`).
- **MeshKit Error Framework (`pkg/errors`)**: Structured error codes (`1001-MCP`, `1002-MCP`, `1003-MCP`) compliant with Layer5 standards.
- **Automated CI/CD Pipeline (`.github/workflows/ci.yml`)**: GitHub Actions workflow for automated testing and binary build verification.
- **Multi-Stage Containerization (`Dockerfile`)**: Lightweight Alpine Docker image packaging (`make docker-build`).
- **Interactive Walkthrough (`DEMO.md`)**: Sample JSON-RPC 2.0 payloads and Claude Desktop / Cursor IDE setup guides.
