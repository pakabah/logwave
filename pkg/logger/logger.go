package logger

import "time"

// Logger interface for structured logging
type Logger interface {
  Debug(message string, additionalLabels map[string]string, additionalMessages map[string]string) error
  Info(message string, additionalLabels map[string]string, additionalMessages map[string]string) error
  Warn(message string, additionalLabels map[string]string, additionalMessages map[string]string) error
  Error(message string, additionalLabels map[string]string, additionalMessages map[string]string) error
  Fatal(message string, additionalLabels map[string]string, additionalMessages map[string]string) error
	Close()
}

// LogEntry defines the structure of a log entry
type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Level     string    `json:"level"`
	Message   string    `json:"message"`
}

// LogMessage holds information about a log to be processed
type LogMessage struct {
	Level, Message string
	Labels         map[string]string
  Messages map[string]string
}
