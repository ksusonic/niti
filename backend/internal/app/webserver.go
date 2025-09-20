package app

import (
	"fmt"
	"time"

	"github.com/ksusonic/niti/backend/internal/api"
	"github.com/ksusonic/niti/backend/internal/api/server"
	"go.uber.org/zap"
)

func (a *App) WebServer() int {
	buildStart := time.Now()
	a.log.Debug("building server deps")

	impl := api.NewAPI(
		a.AuthService(),
		a.UsersRepo(),
		a.SubscriptionsRepo(),
		a.log,
	)

	a.log.Debug("built server deps", zap.Duration("build_took", time.Since(buildStart)))

	engine := server.NewGinServer(
		impl,
		a.AuthService(),
		a.config,
		a.log,
	)

	err := engine.Run(fmt.Sprintf(":%d", a.config.ServerPort))
	if err != nil {
		a.log.Error("web serve", zap.Error(err))
		return 1
	}

	return 0
}
