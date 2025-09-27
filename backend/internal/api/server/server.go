package server

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	authmw "github.com/ksusonic/niti/backend/internal/api/server/middleware/auth"
	"github.com/ksusonic/niti/backend/pgk/config"
	"github.com/ksusonic/niti/backend/pgk/genapi"
	"go.uber.org/zap"
)

func NewGinServer(
	strictServer genapi.StrictServerInterface,
	authDeps authmw.AuthDeps,
	cfg config.WebserverConfig,
	log *zap.Logger,
) *gin.Engine {
	h := genapi.NewStrictHandler(strictServer, nil)

	r := gin.New()
	r.Use(ginzap.Ginzap(log, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(log, true))

	if err := r.SetTrustedProxies(cfg.TrustedProxies); err != nil {
		log.Warn("set trusted proxies", zap.Error(err))
	}

	genapi.RegisterHandlersWithOptions(r, h, genapi.GinServerOptions{
		BaseURL: "",
		Middlewares: []genapi.MiddlewareFunc{
			authmw.AuthMw(authDeps),
		},
		ErrorHandler: func(c *gin.Context, err error, statusCode int) {
			log.Error("unhandled error", zap.Error(err), zap.String("path", c.Request.URL.Path))

			c.JSON(statusCode, genapi.Error{
				Message: "unexpected error",
			})
		},
	})

	return r
}
