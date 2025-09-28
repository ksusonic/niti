package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/ksusonic/niti/backend/internal/app"
)

func main() {
	_ = godotenv.Overload()

	a := app.New()
	defer a.Close(context.Background())

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	a.WebServer(ctx)
}
