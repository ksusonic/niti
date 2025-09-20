package app

import (
	"fmt"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/ksusonic/niti/backend/internal/api"
	"github.com/ksusonic/niti/backend/pgk/genapi"
	middleware "github.com/oapi-codegen/gin-middleware"
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

	validator, err := middleware.OapiValidatorFromYamlFile("../api/openapi.yaml")
	if err != nil {
		a.Logger().Error("openapi validator middleware", zap.Error(err))
		return 1
	}
	r.Use(validator)

	genapi.RegisterHandlers(r, h)

	err = r.Run(fmt.Sprintf(":%d", a.config.ServerPort))
	if err != nil {
		a.log.Error("web serve", zap.Error(err))
		return 1
	}

	return 0
}
