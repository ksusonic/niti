package main

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/ksusonic/niti/backend/internal/app"
	"go.uber.org/zap"
)

func main() {
	_ = godotenv.Overload()

	app := app.New()
	defer app.Close(context.Background())

	bot := app.BotService()

	err := bot.Manage(context.Background())
	if err != nil {
		app.Log.Fatal("manage bot", zap.Error(err))
	}
}
