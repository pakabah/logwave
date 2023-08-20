package logger_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/pakabah/logwave/pkg/logger"
)

func TestLogLevels(t *testing.T) {
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
		{"Warn", "Test warn message", nil, nil, nil},
		{"Error", "Test error message", nil, nil, nil},
		{"Fatal", "Test fatal message", nil, nil, nil},
	}

	for _, s := range scenarios {
    var err error
    switch s.level {
    case "Debug":
      err = lokiLogger.Debug(s.message, s.additionalLabels, s.additionalMsgs)
    case "Info":
      err = lokiLogger.Info(s.message, s.additionalLabels, s.additionalMsgs)
    case "Warn":
      err = lokiLogger.Warn(s.message, s.additionalLabels, s.additionalMsgs)
    case "Error":
      err = lokiLogger.Error(s.message, s.additionalLabels, s.additionalMsgs)
    case "Fatal":
      err = lokiLogger.Fatal(s.message, s.additionalLabels, s.additionalMsgs)
    default:
      t.Errorf("Unknown log level: %s", s.level)
      continue
  }

    if err != s.expectedError {
      t.Errorf("For level %s and message %s, expected error %v but got %v", s.level, s.message, s.expectedError, err)
    }

    time.Sleep(50 * time.Millisecond)
  }

	lokiLogger.Close()
}

// TestSendToLoki_Success tests the successful sending of a log message to Loki.
func TestSendToLoki_Success(t *testing.T) {
	// Creating a mock Loki server that responds with a 200 OK status
	mockLokiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer mockLokiServer.Close()

	// Setting the Loki URL to the mock Loki server's URL
	os.Setenv("LOKI_URL", mockLokiServer.URL)
	config := logger.LoadLokiConfig()
	lokiLogger := logger.NewLokiLogger(config)

	// Sending a test log message
	err := lokiLogger.Debug("Testing http", nil, nil)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	lokiLogger.Close()
}
