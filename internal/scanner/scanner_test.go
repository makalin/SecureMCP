package scanner

import (
	"testing"
)

func TestNewScanner(t *testing.T) {
	scanner := NewScanner()
	if scanner == nil {
		t.Error("NewScanner returned nil")
	}
	if scanner.client == nil {
		t.Error("Scanner client is nil")
	}
}

func TestScan(t *testing.T) {
	scanner := NewScanner()
	results, err := scanner.Scan("http://example.com")
	if err != nil {
		t.Errorf("Scan returned error: %v", err)
	}
	if results == nil {
		t.Error("Scan results is nil")
	}
}

func TestCheckOAuthTokenStorage(t *testing.T) {
	scanner := NewScanner()
	err := scanner.checkOAuthTokenStorage("http://example.com")
	if err != nil {
		t.Errorf("checkOAuthTokenStorage returned error: %v", err)
	}
}

func TestCheckPromptInjection(t *testing.T) {
	scanner := NewScanner()
	err := scanner.checkPromptInjection("http://example.com")
	if err != nil {
		t.Errorf("checkPromptInjection returned error: %v", err)
	}
}

func TestCheckServerAuthentication(t *testing.T) {
	scanner := NewScanner()
	err := scanner.checkServerAuthentication("http://example.com")
	if err != nil {
		t.Errorf("checkServerAuthentication returned error: %v", err)
	}
}
