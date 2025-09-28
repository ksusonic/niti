package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/ksusonic/niti/backend/internal/app"
	"go.uber.org/zap"
)

func main() {
	_ = godotenv.Overload()

	a := app.New()
	defer a.Close(context.Background())

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.Config.Webserver.Port),
		Handler: a.WebHandler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			a.Log.Fatal("listen", zap.Error(err))
		}
	}()

	<-ctx.Done()

	a.Log.Info("shutting down")

	if err := srv.Shutdown(ctx); err != nil {
		a.Log.Error("shutdown", zap.Error(err))
	}
}
