package logger

import (
	"fmt"

	"github.com/ksusonic/niti/backend/pgk/config"
	"go.uber.org/zap"
)

type logFormat string

const (
	SimpleFormat logFormat = "simple"
	JSONFormat   logFormat = "json"
)

func New(config config.LoggerConfig) *zap.Logger {
	var zapConfig zap.Config

	switch logFormat(config.LogFormat) {
	case JSONFormat:
		zapConfig = zap.NewProductionConfig()
	case SimpleFormat:
		zapConfig = zap.NewDevelopmentConfig()
	default:
		zapConfig = zap.NewDevelopmentConfig()
	}

	level, err := zap.ParseAtomicLevel(config.LogLevel)
	if err != nil {
		panic(fmt.Errorf("parse log level: %w", err))
	}

	zapConfig.Level = level
	zapConfig.OutputPaths = []string{"stdout"}
	zapConfig.ErrorOutputPaths = []string{"stderr"}

	zapLogger, err := zapConfig.Build()
	if err != nil {
		panic(fmt.Errorf("build zap config: %w", err))
	}

	return zapLogger
}
