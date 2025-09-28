package logger

import (
	"fmt"

	"github.com/ksusonic/niti/backend/pkg/config"
	"go.uber.org/zap"
)

type logFormat string

const (
	SimpleFormat logFormat = "simple"
	JSONFormat   logFormat = "json"
)

func New(config config.LoggerConfig) (*zap.Logger, error) {
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
		return nil, fmt.Errorf("parse log level: %w", err)
	}

	zapConfig.Level = level
	zapConfig.OutputPaths = []string{"stdout"}
	zapConfig.ErrorOutputPaths = []string{"stderr"}

	zapLogger, err := zapConfig.Build()
	if err != nil {
		return nil, fmt.Errorf("build zap config: %w", err)
	}

	return zapLogger, nil
}
