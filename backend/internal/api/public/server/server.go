package server

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/ksusonic/niti/backend/internal/api/public/server/middleware/auth"
	"github.com/ksusonic/niti/backend/pgk/config"
	"github.com/ksusonic/niti/backend/pgk/publicapi"
	"go.uber.org/zap"
)

func NewGinServer(
	strictServer publicapi.StrictServerInterface,
	authDeps auth.AuthDeps,
	cfg config.WebserverConfig,
	log *zap.Logger,
) *gin.Engine {
	h := publicapi.NewStrictHandler(strictServer, nil)

	r := gin.New()
	r.Use(ginzap.Ginzap(log, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(log, true))

	if err := r.SetTrustedProxies(cfg.TrustedProxies); err != nil {
		log.Warn("set trusted proxies", zap.Error(err))
	}

	publicapi.RegisterHandlersWithOptions(r, h, publicapi.GinServerOptions{
		BaseURL: "",
		Middlewares: []publicapi.MiddlewareFunc{
			auth.AuthMw(authDeps),
		},
		ErrorHandler: func(c *gin.Context, err error, statusCode int) {
			log.Error("unhandled error", zap.Error(err), zap.String("path", c.Request.URL.Path))

			c.JSON(statusCode, publicapi.Error{
				Message: "unexpected error",
			})
		},
	})

	return r
}
