package main

import (
	"log"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/ksusonic/niti/backend/internal/api"
	"github.com/ksusonic/niti/backend/internal/auth"
	"github.com/ksusonic/niti/backend/internal/models"
	"github.com/ksusonic/niti/backend/pgk/genapi"
	"github.com/ksusonic/niti/backend/pgk/logger"
	"go.uber.org/zap"
)

func main() {
	cfg, err := models.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	zapLogger := logger.NewFromEnv()
	defer func() { _ = zapLogger.Sync() }()

	authService := auth.NewService(cfg, zapLogger)

	h := genapi.NewStrictHandler(api.NewServer(authService, zapLogger), nil)

	r := gin.New()
	r.Use(ginzap.Ginzap(zapLogger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(zapLogger, true))

	// Set trusted proxies to avoid trusting all by default
	if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
		zapLogger.Warn("Failed to set trusted proxies", zap.Error(err))
	}

	genapi.RegisterHandlers(r, h)

	err = r.Run(":" + cfg.ServerPort)
	if err != nil {
		zapLogger.Error("Failed to start server", zap.Error(err))
	}
}
