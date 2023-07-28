# LogWave 🌊

LogWave is a Go library designed to facilitate seamless logging to Grafana Loki, with plans to expand its capabilities to other logging platforms in the future.

## Features
- Asynchronous logging to Loki using buffered channels.
- Extensible configuration model to fetch Loki configurations.
- Five logging levels: Debug, Info, Warn, Error, Fatal.
- Support for additional labels in log entries.
- Built with scalability and performance in mind.

## Getting Started

### Installation

To integrate LogWave into your project, use the `go get` command:

```bash
go get github.com/pakabah/logwave
```

### Configuration

LogWave is designed to be easily configurable through environment variables. Before using the library, ensure you have the follwing environment variables set:
- LOKI_URL: This is the URL endpoint of your Loki instance where the logs will be sent. E.g., `http://localhost:3100/loki/api/v1/push`
- LOKI_JOB: Represents the job label for your logs. It can be a descriptor of the application or service producing the logs. E.g., `my-web-app`.
- LOGGING_ENABLED: This determines if logging should be enabled. Set to true if you want to enable logging, otherwise set to `false`.

### Example
```bash
export LOKI_URL=http://localhost:3100/loki/api/v1/push
export LOKI_JOB=my-web-app
export LOGGING_ENABLED=true
```

### Usage

1. First, initialize the Loki logger with the desired configuration:

```go
import "github.com/pakabah/logwave/pkg/logger"

config := logger.LoadLokiConfig()
lokiLogger := logger.NewLokiLogger(config)
```

2. Now you can use the logger to send logs at various levels:

```go
lokiLogger.Debug("This is a debug message", map[string]string{"key": "value"})
lokiLogger.Info("This is an info message", nil)
lokiLogger.Warn("This is a warning", nil)
lokiLogger.Error("Oops! An error occurred", nil)
lokiLogger.Fatal("Fatal error encountered", nil)
```

3. When you're done logging, or before your application exits, remember to close the logger:

```go
lokiLogger.Close()
```

## Future Plans
- Support for more logging platforms apart from Loki.
- Advanced features like log batching, retries, and better error handling.

## License
LogWave is licensed under the MIT License. See [LICENSE](LICENSE) for more details.
