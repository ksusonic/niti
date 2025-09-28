package app

import (
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/ksusonic/niti/backend/internal/api"
	"github.com/ksusonic/niti/backend/internal/api/middleware/auth"
	"github.com/ksusonic/niti/backend/pkg/openapi"
	"go.uber.org/zap"
)

func (a *App) WebHandler() http.Handler {
	r := gin.New()
	r.Use(ginzap.Ginzap(a.Log, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(a.Log, true))

	if err := r.SetTrustedProxies(a.Config.Webserver.TrustedProxies); err != nil {
		a.Log.Warn("set trusted proxies", zap.Error(err))
	}

	openapi.RegisterHandlersWithOptions(
		r,
		openapi.NewStrictHandler(api.NewAPI(
			a.AuthService(),
			a.UsersRepo(),
			a.SubscriptionsRepo(),
			nil,
			a.Log,
		), nil),
		openapi.GinServerOptions{
			BaseURL: "",
			Middlewares: []openapi.MiddlewareFunc{
				auth.Mw(a.AuthService()),
			},
			ErrorHandler: func(c *gin.Context, err error, statusCode int) {
				a.Log.Error("unhandled error", zap.Error(err), zap.String("path", c.Request.URL.Path))

				c.JSON(statusCode, openapi.Error{
					Message: "unexpected error",
				})
			},
		},
	)

	return r
}
