package app

import (
	"github.com/ksusonic/niti/backend/internal/services/auth"
	"github.com/ksusonic/niti/backend/internal/services/bot"
)

type services struct {
	authService *auth.Service
	botService  *bot.Service
}

func (a *App) AuthService() *auth.Service {
	if a.services.authService == nil {
		a.services.authService = auth.NewService(
			a.config.AuthConfig,
			a.RefreshTokenRepo(),
			a.log,
		)
	}

	return a.services.authService
}

func (a *App) BotService() *bot.Service {
	if a.services.botService == nil {
		a.services.botService = bot.NewService(
			a.config.Telegram,
			a.log,
		)
	}

	return a.services.botService
}
