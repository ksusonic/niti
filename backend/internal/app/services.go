package app

import "github.com/ksusonic/niti/backend/internal/services/auth"

type services struct {
	authService *auth.Service
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
