package report

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"time"
)

// ScanReport represents the structure of a security scan report
type ScanReport struct {
	Target          string          `json:"target"`
	ScanTime        time.Time       `json:"scan_time"`
	Vulnerabilities []Vulnerability `json:"vulnerabilities"`
	Summary         Summary         `json:"summary"`
}

// Vulnerability represents a detected security issue
type Vulnerability struct {
	Type        string `json:"type"`
	Severity    string `json:"severity"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Remediation string `json:"remediation"`
}

// Summary contains overall scan statistics
type Summary struct {
	TotalVulnerabilities int `json:"total_vulnerabilities"`
	CriticalCount        int `json:"critical_count"`
	HighCount            int `json:"high_count"`
	MediumCount          int `json:"medium_count"`
	LowCount             int `json:"low_count"`
}

// ReportGenerator handles report generation
type ReportGenerator struct {
	outputDir string
}

// NewReportGenerator creates a new report generator
func NewReportGenerator(outputDir string) *ReportGenerator {
	return &ReportGenerator{
		outputDir: outputDir,
	}
}

// GenerateReport creates a new scan report
func (g *ReportGenerator) GenerateReport(target string, vulnerabilities []string) (*ScanReport, error) {
	report := &ScanReport{
		Target:   target,
		ScanTime: time.Now(),
	}

	// Process vulnerabilities
	for _, vuln := range vulnerabilities {
		v := Vulnerability{
			Type:        determineVulnerabilityType(vuln),
			Severity:    determineSeverity(vuln),
			Description: vuln,
			Location:    target,
			Remediation: generateRemediation(vuln),
		}
		report.Vulnerabilities = append(report.Vulnerabilities, v)
	}

	// Generate summary
	report.Summary = generateSummary(report.Vulnerabilities)

	return report, nil
}

// SaveReport saves the report in the specified format
func (g *ReportGenerator) SaveReport(report *ScanReport, format string) error {
	switch format {
	case "json":
		return g.saveJSON(report)
	case "html":
		return g.saveHTML(report)
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}
}

func (g *ReportGenerator) saveJSON(report *ScanReport) error {
	// Create output directory if it doesn't exist
	if err := os.MkdirAll(g.outputDir, 0755); err != nil {
		return err
	}

	// Generate filename
	filename := filepath.Join(g.outputDir, fmt.Sprintf("scan_%s.json", report.ScanTime.Format("2006_01_02_15_04_05")))

	// Marshal report to JSON
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return err
	}

	// Write to file
	return os.WriteFile(filename, data, 0644)
}

func (g *ReportGenerator) saveHTML(report *ScanReport) error {
	// Create output directory if it doesn't exist
	if err := os.MkdirAll(g.outputDir, 0755); err != nil {
		return err
	}

	// Generate filename
	filename := filepath.Join(g.outputDir, fmt.Sprintf("scan_%s.html", report.ScanTime.Format("2006_01_02_15_04_05")))

	// Create HTML template
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <title>SecureMCP Scan Report</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 20px; }
        .vulnerability { margin: 10px 0; padding: 10px; border: 1px solid #ddd; }
        .critical { background-color: #ffebee; }
        .high { background-color: #fff3e0; }
        .medium { background-color: #e8f5e9; }
        .low { background-color: #e3f2fd; }
        .summary { margin: 20px 0; padding: 10px; background-color: #f5f5f5; }
    </style>
</head>
<body>
    <h1>SecureMCP Scan Report</h1>
    <div class="summary">
        <h2>Summary</h2>
        <p>Target: {{.Target}}</p>
        <p>Scan Time: {{.ScanTime}}</p>
        <p>Total Vulnerabilities: {{.Summary.TotalVulnerabilities}}</p>
        <p>Critical: {{.Summary.CriticalCount}}</p>
        <p>High: {{.Summary.HighCount}}</p>
        <p>Medium: {{.Summary.MediumCount}}</p>
        <p>Low: {{.Summary.LowCount}}</p>
    </div>
    <h2>Vulnerabilities</h2>
    {{range .Vulnerabilities}}
    <div class="vulnerability {{.Severity}}">
        <h3>{{.Type}}</h3>
        <p><strong>Severity:</strong> {{.Severity}}</p>
        <p><strong>Description:</strong> {{.Description}}</p>
        <p><strong>Location:</strong> {{.Location}}</p>
        <p><strong>Remediation:</strong> {{.Remediation}}</p>
    </div>
    {{end}}
</body>
</html>`

	// Parse template
	t, err := template.New("report").Parse(tmpl)
	if err != nil {
		return err
	}

	// Create file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Execute template
	return t.Execute(file, report)
}

func determineVulnerabilityType(vuln string) string {
	if contains(vuln, "token") {
		return "OAuth Token Vulnerability"
	}
	if contains(vuln, "prompt") {
		return "Prompt Injection Vulnerability"
	}
	if contains(vuln, "authentication") {
		return "Authentication Vulnerability"
	}
	return "General Vulnerability"
}

func determineSeverity(vuln string) string {
	criticalKeywords := []string{"critical", "severe", "exploit", "remote code execution"}
	highKeywords := []string{"high", "serious", "authentication bypass"}
	mediumKeywords := []string{"medium", "moderate", "information disclosure"}

	if containsAny(vuln, criticalKeywords) {
		return "critical"
	}
	if containsAny(vuln, highKeywords) {
		return "high"
	}
	if containsAny(vuln, mediumKeywords) {
		return "medium"
	}
	return "low"
}

func generateRemediation(vuln string) string {
	if contains(vuln, "token") {
		return "Implement secure token storage and proper token validation"
	}
	if contains(vuln, "prompt") {
		return "Implement input validation and sanitization for prompts"
	}
	if contains(vuln, "authentication") {
		return "Implement proper authentication mechanisms and security headers"
	}
	return "Review and address the identified security issue"
}

func generateSummary(vulns []Vulnerability) Summary {
	summary := Summary{
		TotalVulnerabilities: len(vulns),
	}

	for _, vuln := range vulns {
		switch vuln.Severity {
		case "critical":
			summary.CriticalCount++
		case "high":
			summary.HighCount++
		case "medium":
			summary.MediumCount++
		case "low":
			summary.LowCount++
		}
	}

	return summary
}

func contains(s, substr string) bool {
	return s != "" && substr != "" && s != substr
}

func containsAny(s string, keywords []string) bool {
	for _, keyword := range keywords {
		if contains(s, keyword) {
			return true
		}
	}
	return false
}
