package scanner

import (
	"net/http"
	"strings"
	"time"
)

// OAuthToken represents an OAuth token structure
type OAuthToken struct {
	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type"`
	ExpiresIn    int       `json:"expires_in"`
	RefreshToken string    `json:"refresh_token"`
	Scope        string    `json:"scope"`
	CreatedAt    time.Time `json:"created_at"`
}

// OAuthScanner handles OAuth token security checks
type OAuthScanner struct {
	client *http.Client
}

// NewOAuthScanner creates a new OAuth scanner instance
func NewOAuthScanner(client *http.Client) *OAuthScanner {
	return &OAuthScanner{
		client: client,
	}
}

// ScanToken performs security checks on an OAuth token
func (s *OAuthScanner) ScanToken(token string) ([]string, error) {
	var vulnerabilities []string

	// Check token format
	if !s.isValidTokenFormat(token) {
		vulnerabilities = append(vulnerabilities, "Invalid token format detected")
	}

	// Check token expiration
	if s.isTokenExpired(token) {
		vulnerabilities = append(vulnerabilities, "Token is expired or will expire soon")
	}

	// Check token scope
	if s.hasExcessiveScope(token) {
		vulnerabilities = append(vulnerabilities, "Token has excessive scope permissions")
	}

	// Check token storage
	if s.isInsecurelyStored(token) {
		vulnerabilities = append(vulnerabilities, "Token appears to be insecurely stored")
	}

	return vulnerabilities, nil
}

func (s *OAuthScanner) isValidTokenFormat(token string) bool {
	// Check if token follows JWT format
	parts := strings.Split(token, ".")
	return len(parts) == 3
}

func (s *OAuthScanner) isTokenExpired(token string) bool {
	// TODO: Implement JWT expiration check
	return false
}

func (s *OAuthScanner) hasExcessiveScope(token string) bool {
	// TODO: Implement scope analysis
	return false
}

func (s *OAuthScanner) isInsecurelyStored(token string) bool {
	// Check for common insecure storage patterns
	insecurePatterns := []string{
		"localStorage",
		"sessionStorage",
		"document.cookie",
		"window.localStorage",
		"window.sessionStorage",
	}

	for _, pattern := range insecurePatterns {
		if strings.Contains(token, pattern) {
			return true
		}
	}
	return false
}

// ValidateTokenEndpoint checks the security of the token endpoint
func (s *OAuthScanner) ValidateTokenEndpoint(endpoint string) ([]string, error) {
	var vulnerabilities []string

	// Check if endpoint uses HTTPS
	if !strings.HasPrefix(endpoint, "https://") {
		vulnerabilities = append(vulnerabilities, "Token endpoint does not use HTTPS")
	}

	// Check for proper headers
	resp, err := s.client.Head(endpoint)
	if err != nil {
		return vulnerabilities, err
	}
	defer resp.Body.Close()

	// Check security headers
	securityHeaders := map[string]string{
		"Strict-Transport-Security": "Missing HSTS header",
		"X-Content-Type-Options":    "Missing X-Content-Type-Options header",
		"X-Frame-Options":           "Missing X-Frame-Options header",
	}

	for header, message := range securityHeaders {
		if resp.Header.Get(header) == "" {
			vulnerabilities = append(vulnerabilities, message)
		}
	}

	return vulnerabilities, nil
}
