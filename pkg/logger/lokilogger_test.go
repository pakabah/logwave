package logger_test

import (
  "os"
  "testing"
  "time"

  "github.com/pakabah/logwave/pkg/logger"
)

func TestDebug(t *testing.T) {
  os.Setenv("LOKI_URL", "http://example.com")
  os.Setenv("LOKI_JOB", "test-job")
  os.Setenv("LOGGING_ENABLED", "true")
  config := logger.LoadLokiConfig()
  lokiLogger := logger.NewLokiLogger(config)

  scenarios := []struct {
    level            string
    message          string
    additionalLabels map[string]string
    additionalMsgs   map[string]string
    expectedError    error
  }{
    {"Debug", "Test debug message", nil, nil, nil},
    {"Info", "Test info message", map[string]string{"labelKey": "labelValue"}, nil, nil},
  }

  for _, s := range scenarios {
    err := lokiLogger.Debug(s.message, s.additionalLabels, s.additionalMsgs)

    if err != s.expectedError {
      t.Errorf("For level %s and message %s, expected error %v but got %v", s.level, s.message, s.expectedError, err)
    }

    time.Sleep(50 * time.Millisecond)
  }

  lokiLogger.Close()
}
