package logger

import (
	"os"
	"strings"
)

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

	labelsStr := os.Getenv("LOKI_LABELS")
	labelsMap := make(map[string]string)
	if labelsStr != "" {
		for _, label := range strings.Split(labelsStr, ",") {
			kv := strings.Split(label, "=")
			if len(kv) == 2 {
				labelsMap[kv[0]] = kv[1]
			}
		}
	}

	return LokiConfig{
		URL:            os.Getenv("LOKI_URL"),
		Job:            os.Getenv("LOKI_JOB"),
		Labels:         labelsMap,
		LoggingEnabled: loggingEnabled,
	}
}
