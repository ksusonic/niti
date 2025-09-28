package main

import (
	"context"

	"github.com/joho/godotenv"
	"github.com/ksusonic/niti/backend/internal/app"
	"go.uber.org/zap"
)

func main() {
	_ = godotenv.Overload()

	a := app.New()
	defer a.Close(context.Background())

	bot := a.BotService()

	err := bot.Manage(context.Background())
	if err != nil {
		a.Log.Fatal("manage bot", zap.Error(err))
	}
}
