# 🔒 Security Policy

## Overview
`meshery-mcp-server` prioritizes data safety and credential protection. It features a dedicated **Response Boundary Secret Sanitizer (`pkg/security`)** designed to prevent token, password, and KubeConfig leakage into LLM context windows or error log streams.

---

## 🛡️ Response Boundary Security Guarantees

1. **Immutability**: Original caller data structures are deep-copied and never mutated during sanitization.
2. **Recursive Redaction**: Deeply nested maps and slice arrays are scrubbed recursively.
3. **Precision Matching**: Key matching is exact and case-insensitive to prevent false-positive over-redaction of non-sensitive fields.
4. **Error Path Scrubbing**: Unstreamed error messages and session cookies are sanitized before logging.

---

## 🚨 Reporting a Vulnerability

If you discover a security vulnerability or credential leak issue in `meshery-mcp-server`, please do NOT open a public GitHub issue.

Please report vulnerabilities directly via email to:
- **`community@layer5.io`**
- **`peaushpaul99@gmail.com`**

We will acknowledge reports within 24 hours and coordinate a security fix release promptly.
