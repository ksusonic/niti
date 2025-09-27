package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ksusonic/niti/backend/internal/api"
	"github.com/ksusonic/niti/backend/internal/api/server"
	"go.uber.org/zap"
)

func (a *App) WebServer(ctx context.Context) {
	buildStart := time.Now()
	a.log.Debug("building server deps")

	impl := api.NewAPI(
		a.AuthService(),
		a.UsersRepo(),
		a.SubscriptionsRepo(),
		nil,
		a.log,
	)

	a.log.Debug("built server deps", zap.Duration("build_took", time.Since(buildStart)))

	engine := server.NewGinServer(
		impl,
		a.AuthService(),
		a.config.Webserver,
		a.log,
	)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.Webserver.Port),
		Handler: engine.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.log.Fatal("listen", zap.Error(err))
		}
	}()

	<-ctx.Done()

	a.log.Info("shutting down")

	if err := srv.Shutdown(ctx); err != nil {
		a.log.Error("shutdown", zap.Error(err))
	}
}
