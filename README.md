# SecureMCP

**SecureMCP** is a comprehensive security auditing tool designed to detect vulnerabilities and misconfigurations in applications using the [Model Context Protocol (MCP)](https://modelcontextprotocol.io/introduction). It proactively identifies threats like OAuth token leakage, prompt injection vulnerabilities, rogue MCP servers, and tool poisoning attacks.

---

## 🛡️ Features

### OAuth Token Scanner
- Token format validation and security checks
- Expiration and scope analysis
- Storage security assessment
- Token endpoint validation
- JWT token analysis

### Prompt Injection Tester
- Multiple injection payload types
- Various injection positions testing
- Response analysis
- System prompt override detection
- Role confusion attack detection

### Authentication & Server Integrity Check
- SSL/TLS configuration validation
- Authentication method testing
- Security header verification
- Server security assessment
- HSTS and CSP validation

### Report Generation
- HTML and JSON report formats
- Vulnerability classification
- Severity assessment
- Remediation suggestions
- Summary statistics

---

## 👨‍💻 Who Should Use SecureMCP?
- AI Developers integrating MCP in applications
- Security teams securing AI model interactions
- DevSecOps engineers embedding MCP in CI/CD pipelines
- Researchers studying AI model vulnerabilities
- Security auditors assessing MCP implementations

---

## 🚀 Getting Started

### Prerequisites
- Go 1.21+
- Docker (optional, for containerized deployment)
- Node.js (for dashboard UI)

### Installation

#### From Source
```bash
git clone https://github.com/makalin/SecureMCP.git
cd SecureMCP
make build
```

#### Using Docker
```bash
docker pull makalin/SecureMCP
```

### Basic Usage

#### Command Line
```bash
# Basic scan
./securemcp scan --target https://your-mcp-server.com

# Scan with specific options
./securemcp scan --target https://your-mcp-server.com \
    --scan-oauth \
    --scan-prompt-injection \
    --scan-authentication \
    --timeout 30s

# Generate HTML report
./securemcp scan --target https://your-mcp-server.com --report html

# Generate JSON report
./securemcp scan --target https://your-mcp-server.com --report json
```

#### Programmatic Usage
```go
import "github.com/makalin/SecureMCP/internal/scanner"

// Create scanner instance
scanner := scanner.NewScanner()

// Basic scan
results, err := scanner.Scan("https://your-mcp-server.com")

// Scan with options
options := &scanner.ScanOptions{
    ScanOAuth:           true,
    ScanPromptInjection: true,
    ScanAuthentication:  true,
    TestPrompt:          "your test prompt",
    Timeout:             30 * time.Second,
}
results, err := scanner.ScanWithOptions(target, options)
```

### Report Generation
```go
import "github.com/makalin/SecureMCP/internal/report"

// Create report generator
generator := report.NewReportGenerator("reports")

// Generate report
report, err := generator.GenerateReport(target, results)

// Save as HTML
err = generator.SaveReport(report, "html")

// Save as JSON
err = generator.SaveReport(report, "json")
```

---

## 📊 Example Output

### Command Line
```bash
$ ./securemcp scan --target https://example-mcp-server.com
[+] Scanning Target: https://example-mcp-server.com
[!] Token storage vulnerability detected
[!] Prompt Injection vulnerability found in tool 'AutoSummary'
[!] Insecure authentication method detected
[+] Report saved to /reports/scan_2024_03_14_15_30_45.html
```

### HTML Report
The HTML report includes:
- Summary statistics
- Vulnerability details
- Severity levels
- Remediation suggestions
- Scan metadata

### JSON Report
```json
{
  "target": "https://example-mcp-server.com",
  "scan_time": "2024-03-14T15:30:45Z",
  "vulnerabilities": [
    {
      "type": "OAuth Token Vulnerability",
      "severity": "high",
      "description": "Token storage vulnerability detected",
      "location": "https://example-mcp-server.com",
      "remediation": "Implement secure token storage and proper token validation"
    }
  ],
  "summary": {
    "total_vulnerabilities": 3,
    "critical_count": 0,
    "high_count": 1,
    "medium_count": 1,
    "low_count": 1
  }
}
```

---

## 🛠️ Development

### Project Structure
```
SecureMCP/
├── cmd/
│   └── securemcp/        # Command-line interface
├── internal/
│   ├── scanner/          # Core scanning functionality
│   │   ├── oauth.go      # OAuth token scanning
│   │   ├── prompt.go     # Prompt injection testing
│   │   ├── auth.go       # Authentication checks
│   │   └── scanner.go    # Main scanner implementation
│   └── report/           # Report generation
├── config/               # Configuration management
├── Dockerfile           # Container configuration
└── Makefile            # Build and development tasks
```

### Building
```bash
# Build binary
make build

# Run tests
make test

# Build Docker image
make docker-build

# Run in Docker
make docker-run
```

---

## 📢 Contributing
Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## 🚀 License
[MIT License](LICENSE)

---

## 🌐 Links
- [GitHub Issues](https://github.com/makalin/SecureMCP/issues)
- [MCP Specification](https://modelcontextprotocol.io/introduction)
- [Documentation](https://github.com/makalin/SecureMCP/wiki)

---

Protect your MCP applications before they get exploited. 💪 **Use SecureMCP!**
