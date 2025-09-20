package auth

import (
	"time"

	"github.com/ksusonic/niti/backend/pgk/config"
	"go.uber.org/zap"
)

type Service struct {
	token         string
	expiresIn     time.Duration
	accessSecret  []byte
	refreshSecret []byte
	accessTTL     time.Duration
	refreshTTL    time.Duration
	logger        *zap.Logger
}

func NewService(cfg config.AuthConfig, logger *zap.Logger) *Service {
	if cfg.AccessSecret == "" {
		logger.Fatal("ACCESS_SECRET required")
	}
	if cfg.RefreshSecret == "" {
		logger.Fatal("REFRESH_SECRET required")
	}
	if cfg.TelegramToken == "" {
		logger.Fatal("TELEGRAM_TOKEN required")
	}

	return &Service{
		token:         cfg.TelegramToken,
		expiresIn:     cfg.TelegramExpiresIn,
		accessSecret:  []byte(cfg.AccessSecret),
		refreshSecret: []byte(cfg.RefreshSecret),
		accessTTL:     cfg.AccessTTL,
		refreshTTL:    cfg.RefreshTTL,
		logger:        logger,
	}
}
