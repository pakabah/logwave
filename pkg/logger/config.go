package logger

import (
	"log"
	"os"
	"strings"
)

// LokiConfig defines configurations for Loki logging
// LokiConfig holds the configuration for connecting to Grafana Loki.
// It includes the URL, job name, labels, and a flag to enable or disable logging./ LokiConfig defines configurations for Loki logging
type LokiConfig struct {
	URL            string
	Job            string
	Labels         map[string]string
	LoggingEnabled bool
}

// LoadLokiConfig loads Loki configurations from the environment
// LoadLokiConfig loads the Loki configuration from environment variables.
// It reads LOKI_URL, LOKI_JOB, LOKI_LABELS, and LOGGING_ENABLED and returns a LokiConfig object.
func LoadLokiConfig() LokiConfig {
	loggingEnabled := os.Getenv("LOGGING_ENABLED") == "true"

	labelsStr := os.Getenv("LOKI_LABELS")
	labelsMap := make(map[string]string)
	if labelsStr != "" {
		for _, label := range strings.Split(labelsStr, ",") {
			kv := strings.Split(label, "=")
			if len(kv) == 2 {
				labelsMap[kv[0]] = kv[1]
			} else {
        log.Printf("Warning: Malformed label detected: '%s'. Expected format: 'key=value'", label)
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
