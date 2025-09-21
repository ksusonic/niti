package app

import (
	"context"

	refreshtoken "github.com/ksusonic/niti/backend/internal/storage/repository/refresh_token"
	"github.com/ksusonic/niti/backend/internal/storage/repository/subscriptions"
	"github.com/ksusonic/niti/backend/internal/storage/repository/users"
)

type repositories struct {
	usersRepo         *users.Repository
	subscriptionsRepo *subscriptions.Repository
	refreshTokenRepo  *refreshtoken.Repository
}

func (a *App) UsersRepo() *users.Repository {
	if a.repositories.usersRepo == nil {
		a.repositories.usersRepo = users.New(a.postgresPool(context.TODO()).Pool)
	}

	return a.repositories.usersRepo
}

func (a *App) SubscriptionsRepo() *subscriptions.Repository {
	if a.repositories.subscriptionsRepo == nil {
		a.repositories.subscriptionsRepo = subscriptions.New(a.postgresPool(context.TODO()).Pool)
	}

	return a.repositories.subscriptionsRepo
}

func (a *App) RefreshTokenRepo() *refreshtoken.Repository {
	if a.repositories.refreshTokenRepo == nil {
		a.repositories.refreshTokenRepo = refreshtoken.New(a.postgresPool(context.TODO()).Pool)
	}

	return a.repositories.refreshTokenRepo
}
