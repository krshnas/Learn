package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LoggerConfig holds the configuration for initializing a logger.
type LoggerConfig struct {
	Level            zapcore.Level
	Development      bool
	OutputPaths      []string
	ErrorOutputPaths []string
	InitialFields    map[string]interface{}
}

// DefaultLoggerConfig provides the default configuration for the logger.
func DefaultLoggerConfig() LoggerConfig {
	return LoggerConfig{
		Level:            zap.InfoLevel,
		Development:      false,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}
}

// InitializeLogger creates and returns a zap.Logger based on the provided config.
func InitializeLogger(cfg LoggerConfig) (*zap.Logger, error) {
	// Create an encoder config for structured JSON logging
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	// Set the logger's log level
	level := zap.NewAtomicLevelAt(cfg.Level)

	// Create the zap configuration using the provided values
	zapConfig := zap.Config{
		Level:             level,
		Development:       cfg.Development,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths:       cfg.OutputPaths,
		ErrorOutputPaths:  cfg.ErrorOutputPaths,
		InitialFields:     cfg.InitialFields,
	}

	// Build and return the logger
	logger, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
