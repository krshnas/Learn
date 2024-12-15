package main

import (
	"github.com/krishna/go/learn/cli/cmd/stringer"
	"github.com/krishna/go/learn/cli/pkg/log"
	"go.uber.org/zap"
)

// setupLogger initializes and returns a logger based on the provided configuration.
func setupLogger() (*zap.Logger, error) {
	// Get the default logger configuration
	cfg := log.DefaultLoggerConfig()

	// Modify config for development environment
	cfg.Development = true
	cfg.Level = zap.DebugLevel

	// Initialize and return the logger
	return log.InitializeLogger(cfg)
}

func main() {
	// Initialize the logger
	logger, err := setupLogger()
	if err != nil {
		// Use structured logging with zap
		logger.Error("Error initializing logger", zap.Error(err))
		// Optionally, exit the program after logging the error
		return
	}

	// Ensure the logger is flushed before the application exits
	defer logger.Sync()

	// Execute the stringer command
	stringer.Execute(logger)
}
