package config

import (
	"os"
	"strconv"
)

// Config holds the application configuration
type Config struct {
	ServerPort    int
	ScanTimeout   int
	ReportDir     string
	LogLevel      string
	EnableMetrics bool
}

// Load loads the configuration from environment variables
func Load() *Config {
	port, _ := strconv.Atoi(getEnv("SERVER_PORT", "8080"))
	timeout, _ := strconv.Atoi(getEnv("SCAN_TIMEOUT", "30"))
	enableMetrics, _ := strconv.ParseBool(getEnv("ENABLE_METRICS", "false"))

	return &Config{
		ServerPort:    port,
		ScanTimeout:   timeout,
		ReportDir:     getEnv("REPORT_DIR", "reports"),
		LogLevel:      getEnv("LOG_LEVEL", "info"),
		EnableMetrics: enableMetrics,
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
