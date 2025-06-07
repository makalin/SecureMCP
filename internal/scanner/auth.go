package scanner

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// AuthScanner handles server authentication security checks
type AuthScanner struct {
	client *http.Client
}

// NewAuthScanner creates a new authentication scanner instance
func NewAuthScanner(client *http.Client) *AuthScanner {
	return &AuthScanner{
		client: client,
	}
}

// ScanAuth performs authentication security checks
func (s *AuthScanner) ScanAuth(endpoint string) ([]string, error) {
	var vulnerabilities []string

	// Check SSL/TLS configuration
	sslVulns, err := s.checkSSLConfig(endpoint)
	if err != nil {
		return vulnerabilities, err
	}
	vulnerabilities = append(vulnerabilities, sslVulns...)

	// Check authentication methods
	authVulns, err := s.checkAuthMethods(endpoint)
	if err != nil {
		return vulnerabilities, err
	}
	vulnerabilities = append(vulnerabilities, authVulns...)

	// Check security headers
	headerVulns, err := s.checkSecurityHeaders(endpoint)
	if err != nil {
		return vulnerabilities, err
	}
	vulnerabilities = append(vulnerabilities, headerVulns...)

	return vulnerabilities, nil
}

func (s *AuthScanner) checkSSLConfig(endpoint string) ([]string, error) {
	var vulnerabilities []string

	// Create custom transport for SSL checks
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // We want to check the certificate
		},
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}

	// Make request to check SSL
	resp, err := client.Get(endpoint)
	if err != nil {
		if strings.Contains(err.Error(), "certificate") {
			vulnerabilities = append(vulnerabilities, "Invalid SSL certificate")
		}
		return vulnerabilities, nil
	}
	defer resp.Body.Close()

	// Check TLS version
	if resp.TLS != nil {
		if resp.TLS.Version < tls.VersionTLS12 {
			vulnerabilities = append(vulnerabilities, "Outdated TLS version")
		}
	}

	return vulnerabilities, nil
}

func (s *AuthScanner) checkAuthMethods(endpoint string) ([]string, error) {
	var vulnerabilities []string

	// Test common authentication methods
	authMethods := []string{
		"Basic",
		"Bearer",
		"Digest",
		"OAuth",
		"JWT",
	}

	for _, method := range authMethods {
		req, err := http.NewRequest("GET", endpoint, nil)
		if err != nil {
			continue
		}

		// Add authentication header
		req.Header.Add("Authorization", fmt.Sprintf("%s test", method))

		resp, err := s.client.Do(req)
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		// Check response for authentication method support
		if resp.StatusCode == http.StatusUnauthorized {
			authHeader := resp.Header.Get("WWW-Authenticate")
			if authHeader == "" {
				vulnerabilities = append(vulnerabilities,
					fmt.Sprintf("Missing WWW-Authenticate header for %s auth", method))
			}
		}
	}

	return vulnerabilities, nil
}

func (s *AuthScanner) checkSecurityHeaders(endpoint string) ([]string, error) {
	var vulnerabilities []string

	resp, err := s.client.Get(endpoint)
	if err != nil {
		return vulnerabilities, err
	}
	defer resp.Body.Close()

	// Required security headers
	requiredHeaders := map[string]string{
		"Strict-Transport-Security": "Missing HSTS header",
		"X-Content-Type-Options":    "Missing X-Content-Type-Options header",
		"X-Frame-Options":           "Missing X-Frame-Options header",
		"X-XSS-Protection":          "Missing X-XSS-Protection header",
		"Content-Security-Policy":   "Missing Content-Security-Policy header",
	}

	for header, message := range requiredHeaders {
		if resp.Header.Get(header) == "" {
			vulnerabilities = append(vulnerabilities, message)
		}
	}

	return vulnerabilities, nil
}
