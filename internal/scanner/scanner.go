package scanner

import (
	"fmt"
	"net/http"
	"time"
)

// Scanner represents the main scanning functionality
type Scanner struct {
	client        *http.Client
	oauthScanner  *OAuthScanner
	promptScanner *PromptScanner
	authScanner   *AuthScanner
}

// NewScanner creates a new Scanner instance
func NewScanner() *Scanner {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	return &Scanner{
		client:        client,
		oauthScanner:  NewOAuthScanner(client),
		promptScanner: NewPromptScanner(client),
		authScanner:   NewAuthScanner(client),
	}
}

// Scan performs a security scan on the target MCP server
func (s *Scanner) Scan(target string) ([]string, error) {
	var results []string

	// Check OAuth token storage
	oauthVulns, err := s.oauthScanner.ScanToken(target)
	if err != nil {
		return results, fmt.Errorf("OAuth scan failed: %v", err)
	}
	results = append(results, oauthVulns...)

	// Check prompt injection vulnerabilities
	promptVulns, err := s.promptScanner.ScanPrompt(target, "test prompt")
	if err != nil {
		return results, fmt.Errorf("Prompt injection scan failed: %v", err)
	}
	results = append(results, promptVulns...)

	// Check server authentication
	authVulns, err := s.authScanner.ScanAuth(target)
	if err != nil {
		return results, fmt.Errorf("Authentication scan failed: %v", err)
	}
	results = append(results, authVulns...)

	return results, nil
}

// ScanWithOptions performs a security scan with specific options
func (s *Scanner) ScanWithOptions(target string, options *ScanOptions) ([]string, error) {
	var results []string

	if options.ScanOAuth {
		oauthVulns, err := s.oauthScanner.ScanToken(target)
		if err != nil {
			return results, fmt.Errorf("OAuth scan failed: %v", err)
		}
		results = append(results, oauthVulns...)
	}

	if options.ScanPromptInjection {
		promptVulns, err := s.promptScanner.ScanPrompt(target, options.TestPrompt)
		if err != nil {
			return results, fmt.Errorf("Prompt injection scan failed: %v", err)
		}
		results = append(results, promptVulns...)
	}

	if options.ScanAuthentication {
		authVulns, err := s.authScanner.ScanAuth(target)
		if err != nil {
			return results, fmt.Errorf("Authentication scan failed: %v", err)
		}
		results = append(results, authVulns...)
	}

	return results, nil
}

// ScanOptions represents options for the security scan
type ScanOptions struct {
	ScanOAuth           bool
	ScanPromptInjection bool
	ScanAuthentication  bool
	TestPrompt          string
	Timeout             time.Duration
}

// DefaultScanOptions returns default scan options
func DefaultScanOptions() *ScanOptions {
	return &ScanOptions{
		ScanOAuth:           true,
		ScanPromptInjection: true,
		ScanAuthentication:  true,
		TestPrompt:          "test prompt",
		Timeout:             30 * time.Second,
	}
}
