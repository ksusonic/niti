package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ksusonic/niti/backend/internal/api/public"
	"github.com/ksusonic/niti/backend/internal/api/public/server"
	"go.uber.org/zap"
)

func (a *App) WebServer(ctx context.Context) {
	buildStart := time.Now()
	a.Log.Debug("building server deps")

	impl := public.NewAPI(
		a.AuthService(),
		a.UsersRepo(),
		a.SubscriptionsRepo(),
		nil,
		a.Log,
	)

	a.Log.Debug("built server deps", zap.Duration("build_took", time.Since(buildStart)))

	engine := server.NewGinServer(
		impl,
		a.AuthService(),
		a.Config.Webserver,
		a.Log,
	)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.Config.Webserver.Port),
		Handler: engine.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			a.Log.Fatal("listen", zap.Error(err))
		}
	}()

	<-ctx.Done()

	a.Log.Info("shutting down")

	if err := srv.Shutdown(ctx); err != nil {
		a.Log.Error("shutdown", zap.Error(err))
	}
}
