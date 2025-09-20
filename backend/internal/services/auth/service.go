package auth

import (
	"github.com/ksusonic/niti/backend/pgk/config"
	"go.uber.org/zap"
)

type Service struct {
	cfg              config.AuthConfig
	refreshTokenRepo refreshTokenRepo
	logger           *zap.Logger
}

func NewService(
	cfg config.AuthConfig,
	refreshTokenRepo refreshTokenRepo,
	logger *zap.Logger,
) *Service {
	if cfg.AccessSecret == "" {
		logger.Fatal("ACCESS_SECRET required")
	}
	if cfg.RefreshSecret == "" {
		logger.Fatal("REFRESH_SECRET required")
	}
	if cfg.TelegramToken == "" {
		logger.Fatal("TELEGRAM_BOT_TOKEN required")
	}

	return &Service{
		cfg:              cfg,
		refreshTokenRepo: refreshTokenRepo,
		logger:           logger,
	}
}
