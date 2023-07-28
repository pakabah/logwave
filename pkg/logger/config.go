package logger

import "os"

// LokiConfig defines configurations for Loki logging
type LokiConfig struct {
    URL            string
    Job            string
    Labels         map[string]string
    LoggingEnabled bool
}

// LoadLokiConfig loads Loki configurations from the environment
func LoadLokiConfig() LokiConfig {
    loggingEnabled := os.Getenv("LOGGING_ENABLED") == "true"

    return LokiConfig{
        URL:           os.Getenv("LOKI_URL"),
        Job:           os.Getenv("LOKI_JOB"),
        Labels:        map[string]string{},
        LoggingEnabled: loggingEnabled,
    }
}

