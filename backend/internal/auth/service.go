package auth

import (
	"time"

	"github.com/ksusonic/niti/backend/internal/models"
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

func NewService(cfg *models.Config, logger *zap.Logger) *Service {
	return &Service{
		token:         cfg.TelegramToken,
		expiresIn:     cfg.TelegramExpiresIn,
		accessSecret:  cfg.AccessSecret,
		refreshSecret: cfg.RefreshSecret,
		accessTTL:     cfg.AccessTTL,
		refreshTTL:    cfg.RefreshTTL,
		logger:        logger,
	}
}
