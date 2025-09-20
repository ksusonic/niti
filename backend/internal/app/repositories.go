package app

import (
	"context"

	refreshtoken "github.com/ksusonic/niti/backend/internal/storage/repository/refresh_token"
)

type repositories struct {
	refreshTokenRepo *refreshtoken.Repository
}

func (a *App) RefreshTokenRepo() *refreshtoken.Repository {
	if a.repositories.refreshTokenRepo == nil {
		a.repositories.refreshTokenRepo = refreshtoken.New(a.postgresPool(context.TODO()).Pool)
	}

	return a.repositories.refreshTokenRepo
}
