# SecureMCP

**SecureMCP** is a security auditing tool designed to detect vulnerabilities and misconfigurations in applications using the [Model Context Protocol (MCP)](https://modelcontextprotocol.io/introduction). It proactively identifies threats like OAuth token leakage, prompt injection vulnerabilities, rogue MCP servers, and tool poisoning attacks.

---

## 🛡️ Features

- **OAuth Token Scanner**
  - Detects insecure token storage and access issues.

- **Prompt Injection Tester**
  - Simulates malicious prompt payloads to check LLM vulnerability.

- **Authentication & Server Integrity Check**
  - Validates MCP servers, SSL certificates, and proper authentication mechanisms.

- **Tool Poisoning Detector**
  - Analyzes MCP tool definitions for suspicious updates or malicious behavior.

- **Supply Chain Audit**
  - Inspects external dependencies and highlights potential risks.

- **Continuous Monitoring**
  - Provides real-time alerts for new threats or unauthorized tool updates.

- **Penetration Testing Toolkit (Pro Mode)**
  - Custom attack simulations against your MCP environment.

- **Risk Dashboard & Reports**
  - Generates risk scores and detailed vulnerability reports in PDF/JSON.

---

## 👨‍💻 Who Should Use SecureMCP?
- AI Developers integrating MCP in applications.
- Security teams securing AI model interactions.
- DevSecOps engineers embedding MCP in CI/CD pipelines.
- Researchers studying AI model vulnerabilities.

---

## 🚀 Getting Started

### Prerequisites
- Go 1.21+ or Rust (depending on chosen build)
- Docker (optional, for containerized deployment)
- Node.js (for dashboard UI)

### Installation
```bash
git clone https://github.com/makalin/SecureMCP.git
cd SecureMCP
make build
```

Or use Docker:
```bash
docker pull makalin/SecureMCP
```

### Running
```bash
./securemcp scan --target https://your-mcp-server.com
```

### Output
- HTML Dashboard
- JSON report
- CVE-style vulnerability summaries

---

## 📈 Example
```bash
securemcp scan --target https://example-mcp-server.com --report html
```

Sample output:
```
[+] Scanning Target: https://example-mcp-server.com
[!] Token storage vulnerability detected.
[!] Prompt Injection vulnerability found in tool 'AutoSummary'.
[!] Insecure authentication method detected.
Report saved to /reports/scan_2025_06_05.html
```

---

## 🎓 Documentation
- [MCP Specification](https://modelcontextprotocol.io/introduction)
- [SecureMCP Docs](https://github.com/makalin/SecureMCP/)

---

## 📢 Contributing
Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

---

## 👾 Security
If you find a security vulnerability, please report it privately to security@yourdomain.com.

---

## 🚀 License
[MIT License](LICENSE)

---

## 🌐 Links
- [GitHub Issues](https://github.com/makalin/SecureMCP/issues)

---

Protect your MCP applications before they get exploited. 💪  **Use SecureMCP!**
