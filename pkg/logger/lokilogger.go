package logger

import (
  "bytes"
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "strconv"
  "time"
)

const (
  DEBUG = "DEBUG"
  INFO  = "INFO"
  WARN  = "WARN"
  ERROR = "ERROR"
  FATAL = "FATAL"
)

// LokiLogger implements the Logger interface for Loki
type LokiLogger struct {
  Config     LokiConfig
  logChannel chan LogMessage
}

// NewLokiLogger creates a new Loki logger instance
func NewLokiLogger(config LokiConfig) Logger {
  logger := &LokiLogger{
    Config:     config,
    logChannel: make(chan LogMessage, 1000), // Buffer size of 1000, adjust as needed
  }
  go logger.processLogs()
  return logger
}

// processLogs processes log messages asynchronously
func (l *LokiLogger) processLogs() {
  for logMessage := range l.logChannel {
    err := l.sendToLoki(logMessage.Level, logMessage.Message, logMessage.Labels)
    if err != nil {
      log.Printf("Error sending log to Loki: %v", err)
    }
  }
}

// sendToLoki sends log messages to Loki
func (l *LokiLogger) sendToLoki(level, message string, additionalLabels map[string]string) error {
  config := l.Config

  logLabels := map[string]string{
    "job": config.Job,
  }

  for k, v := range config.Labels {
    logLabels[k] = v
  }
  for k, v := range additionalLabels {
    logLabels[k] = v
  }

  entry := LogEntry{
    Timestamp: time.Now(),
    Level:     level,
    Message:   message,
  }

  serializedEntry, err := json.Marshal(entry)
  if err != nil {
    return err
  }

  currentTimeNano := strconv.FormatInt(entry.Timestamp.UnixNano(), 10)
  values := [][2]string{
    {
      currentTimeNano,
      string(serializedEntry),
    },
  }

  logEntry := map[string]interface{}{
    "streams": []map[string]interface{}{
      {
        "stream": logLabels,
        "values": values,
      },
    },
  }

  data, err := json.Marshal(logEntry)
  if err != nil {
    return err
  }

  resp, err := http.Post(config.URL, "application/json", bytes.NewBuffer(data))
  if err != nil {
    return err
  }

  defer resp.Body.Close()
  if resp.StatusCode != http.StatusOK {
    return fmt.Errorf("Failed to send log to Loki. HTTP response code: %d", resp.StatusCode)
  }

  return nil
}

// Send is a general function to send logs
func (l *LokiLogger) Send(level, message string, additionalLabels, additionalMessages map[string]string) error {
  if !l.Config.LoggingEnabled {
    return nil
  }

  l.logChannel <- LogMessage{Level: level, Message: message, Labels: additionalLabels, Messages: additionalMessages}
  return nil
}

// Close closes the logger's channel
func (l *LokiLogger) Close() {
  close(l.logChannel)
}

// Debug logs a debug message
func (l *LokiLogger) Debug(message string, additionalLabels, additionalMessages map[string]string) error {
  return l.Send(DEBUG, message, additionalLabels, additionalMessages)
}

// Info logs an informational message
func (l *LokiLogger) Info(message string, additionalLabels, additionalMessages map[string]string) error {
  return l.Send(INFO, message, additionalLabels, additionalMessages)
}

// Warn logs a warning message
func (l *LokiLogger) Warn(message string, additionalLabels, additionalMessages map[string]string) error {
  return l.Send(WARN, message, additionalLabels, additionalMessages)
}

// Error logs an error message
func (l *LokiLogger) Error(message string, additionalLabels, additionalMessages map[string]string) error {
  return l.Send(ERROR, message, additionalLabels, additionalMessages)
}

// Fatal logs a fatal error message
func (l *LokiLogger) Fatal(message string, additionalLabels, additionalMessages map[string]string) error {
  return l.Send(FATAL, message, additionalLabels, additionalMessages)
}
