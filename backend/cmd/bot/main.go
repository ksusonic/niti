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

	app := app.New()
	defer app.Close(context.Background())

	bot := app.BotService()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	bot.Start(ctx)
}
