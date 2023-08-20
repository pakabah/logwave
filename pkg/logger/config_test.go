package logger_test

import (
  "os"
  "testing"

  "github.com/pakabah/logwave/pkg/logger"
)

func TestLoadLokiConfig(t *testing.T) {
  os.Setenv("LOKI_URL", "http://example.com")
  os.Setenv("LOKI_JOB", "test-job")
  os.Setenv("LOKI_LABELS", "key1=value1,key2=value2")
  os.Setenv("LOGGING_ENABLED", "true")

  config := logger.LoadLokiConfig()

  if config.URL != "http://example.com" || config.Job != "test-job" || config.LoggingEnabled != true {
    t.Fatalf("Failed to load configuration correctly")
  }

  if val, ok := config.Labels["key1"]; !ok || val != "value1" {
    t.Errorf("Failed to parse LOKI_LABELS correctly")
  }

  os.Unsetenv("LOKI_LABELS")
  configWithoutLabels := logger.LoadLokiConfig()

  if len(configWithoutLabels.Labels) != 0 {
    t.Errorf("Expected no labels, but got: %v", configWithoutLabels.Labels)
  }
}

func TestLoadLokiConfigLoggingEnabled(t *testing.T) {
  // Setting up the environment
  os.Setenv("LOKI_URL", "http://example.com")
  os.Setenv("LOKI_JOB", "test-job")
  os.Setenv("LOKI_LABELS", "key1=value1,key2=value2")
  os.Setenv("LOGGING_ENABLED", "true")

  config := logger.LoadLokiConfig()

  if config.LoggingEnabled != true {
    t.Fatalf("Failed to load LOGGING_ENABLED correctly when set to true")
  }

  // Test for when LOGGING_ENABLED is not available
  os.Unsetenv("LOGGING_ENABLED")
  configWithoutLoggingEnabled := logger.LoadLokiConfig()

  if configWithoutLoggingEnabled.LoggingEnabled != false {
    t.Errorf("Expected LoggingEnabled to be false when LOGGING_ENABLED is not set, but got: %v", configWithoutLoggingEnabled.LoggingEnabled)
  }
}

func TestMissingEnvironmentVariables(t *testing.T) {
  // Unset environment variables
  os.Unsetenv("LOKI_URL")
  os.Unsetenv("LOKI_JOB")

  config := logger.LoadLokiConfig()

  // Validate specific fields or conditions based on your implementation
  if config.URL == "" {
    t.Errorf("Expected LokiURL to be populated, but it is empty")
  }
  if config.Job == "" {
    t.Errorf("Expected Job to be populated, but it is empty")
  }
}

