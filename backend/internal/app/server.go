package app

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/ksusonic/niti/backend/internal/api"
	"github.com/ksusonic/niti/backend/pgk/genapi"
	"go.uber.org/zap"
)

func (a *App) WebServe() int {
	buildStart := time.Now()
	a.log.Debug("building server deps")

	h := genapi.NewStrictHandler(api.NewServer(
		a.AuthService(),
		a.Logger(),
	), nil)

	a.log.Debug(
		"built server deps",
		zap.Duration("build_took", time.Since(buildStart)),
	)

	r := gin.New()
	r.Use(ginzap.Ginzap(a.Logger(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(a.Logger(), true))

	if err := r.SetTrustedProxies(a.config.TrustedProxies); err != nil {
		a.Logger().Warn("Failed to set trusted proxies", zap.Error(err))
	}

	// healthcheck endpoint
	r.Handle("GET", "/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	genapi.RegisterHandlers(r, h)

	err := r.Run(":" + a.config.ServerPort)
	if err != nil {
		a.log.Error("web serve", zap.Error(err))
		return 1
	}

	return 0
}
