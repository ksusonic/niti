package server

import (
	"fmt"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/ksusonic/niti/backend/internal/api/public/server/middleware/auth"
	"github.com/ksusonic/niti/backend/pkg/config"
	"github.com/ksusonic/niti/backend/pkg/privateapi"
	"github.com/ksusonic/niti/backend/pkg/publicapi"
	"go.uber.org/zap"
)

func NewGinServer(
	publicStrictServer publicapi.StrictServerInterface,
	privateStrictServer privateapi.StrictServerInterface,
	authDeps auth.Deps,
	cfg config.WebserverConfig,
	log *zap.Logger,
) *gin.Engine {
	r := gin.New()
	r.Use(ginzap.Ginzap(log, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(log, true))

	if err := r.SetTrustedProxies(cfg.TrustedProxies); err != nil {
		log.Warn("set trusted proxies", zap.Error(err))
	}

	publicapi.RegisterHandlersWithOptions(
		r,
		publicapi.NewStrictHandler(publicStrictServer, nil),
		publicapi.GinServerOptions{
			BaseURL: "",
			Middlewares: []publicapi.MiddlewareFunc{
				auth.Mw(authDeps),
			},
			ErrorHandler: func(c *gin.Context, err error, statusCode int) {
				log.Error("unhandled error", zap.Error(err), zap.String("path", c.Request.URL.Path))

				c.JSON(statusCode, publicapi.Error{
					Message: "unexpected error",
				})
			},
		},
	)

	privateapi.RegisterHandlersWithOptions(
		r,
		privateapi.NewStrictHandler(privateStrictServer, nil),
		privateapi.GinServerOptions{
			BaseURL:     "/private",
			Middlewares: nil,
			ErrorHandler: func(c *gin.Context, err error, statusCode int) {
				log.Error("unhandled error", zap.Error(err), zap.String("path", c.Request.URL.Path))

				c.JSON(statusCode, privateapi.Error{
					Message: fmt.Sprintf("unexpected error: %v", err),
				})
			},
		},
	)

	return r
}
