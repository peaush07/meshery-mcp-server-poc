# 🤝 Contributing to Meshery MCP Server

Thank you for your interest in contributing to **Meshery MCP Server**! This project follows CNCF and Layer5 open-source community conventions.

---

## 🚀 Development Workflow

### 1. Prerequisites
- **Go 1.22+** installed locally.
- **Git** configured with your signed-off commits (`git commit -s`).
- **Docker** (optional, for container builds).

### 2. Testing & Verification
Before submitting a Pull Request, ensure all unit tests and performance benchmarks pass:

```bash
# Run unit test suite
make test

# Run performance benchmarks
make bench

# Verify binary build
make build
```

---

## 🔒 Code Standards & Security

1. **Response Boundary Sanitization**: All new MCP tool outputs must be passed through `pkg/security.SanitizeMap()` or `pkg/security.SanitizeJSON()` to guarantee sensitive credentials (tokens, KubeConfigs, secrets) are redacted.
2. **MeshKit Errors**: Use structured error definitions from `pkg/errors` so error outputs conform to Meshery error-code conventions.
3. **Clean Code**: Run `gofmt -s -w .` before committing code.

---

## 📄 License & DCO

All contributions are submitted under the **Apache License 2.0** and require Developer Certificate of Origin (DCO) sign-off (`git commit -s`).
