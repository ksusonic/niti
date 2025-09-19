package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Level  LogLevel
	Format LogFormat
}

type LogLevel string

const (
	DebugLevel LogLevel = "debug"
	InfoLevel  LogLevel = "info"
	WarnLevel  LogLevel = "warn"
	ErrorLevel LogLevel = "error"
)

type LogFormat string

const (
	SimpleFormat LogFormat = "simple"
	JSONFormat   LogFormat = "json"
)

// New creates a new zap logger instance based on the provided configuration
func New(config Config) *zap.Logger {
	// Parse log level
	var level zapcore.Level
	switch config.Level {
	case DebugLevel:
		level = zapcore.DebugLevel
	case InfoLevel:
		level = zapcore.InfoLevel
	case WarnLevel:
		level = zapcore.WarnLevel
	case ErrorLevel:
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	var zapConfig zap.Config

	switch config.Format {
	case JSONFormat:
		zapConfig = zap.NewProductionConfig()
	case SimpleFormat:
		zapConfig = zap.NewDevelopmentConfig()
	default:
		zapConfig = zap.NewDevelopmentConfig()
	}

	zapConfig.Level = zap.NewAtomicLevelAt(level)
	zapConfig.OutputPaths = []string{"stdout"}
	zapConfig.ErrorOutputPaths = []string{"stderr"}

	zapLogger, err := zapConfig.Build()
	if err != nil {
		// Fallback to a basic logger if configuration fails
		zapLogger = zap.NewNop()
	}

	return zapLogger
}

// NewFromEnv creates a logger using environment variables
func NewFromEnv() *zap.Logger {
	config := Config{
		Level:  LogLevel(getEnvOrDefault("LOG_LEVEL", "info")),
		Format: LogFormat(getEnvOrDefault("LOG_FORMAT", "simple")),
	}
	return New(config)
}

// Helper function to get environment variable with default
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
