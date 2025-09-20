package server

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/ksusonic/niti/backend/pgk/config"
	"github.com/ksusonic/niti/backend/pgk/genapi"
	"go.uber.org/zap"
)

func NewGinServer(
	strictServer genapi.StrictServerInterface,
	cfg *config.Config,
	log *zap.Logger,
) *gin.Engine {
	h := genapi.NewStrictHandler(strictServer, nil)

	r := gin.New()
	r.Use(ginzap.Ginzap(log, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(log, true))

	if err := r.SetTrustedProxies(cfg.TrustedProxies); err != nil {
		log.Warn("set trusted proxies", zap.Error(err))
	}

	genapi.RegisterHandlers(r, h)

	return r
}
