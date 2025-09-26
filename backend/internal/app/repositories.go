package app

import (
	"context"

	"github.com/ksusonic/niti/backend/internal/storage"
	refreshtoken "github.com/ksusonic/niti/backend/internal/storage/repository/refresh_token"
	"github.com/ksusonic/niti/backend/internal/storage/repository/subscriptions"
	"github.com/ksusonic/niti/backend/internal/storage/repository/users"
	"go.uber.org/zap"
)

type repositories struct {
	usersRepo         *users.Repository
	subscriptionsRepo *subscriptions.Repository
	refreshTokenRepo  *refreshtoken.Repository
}

func (a *App) postgresPool(ctx context.Context) *storage.Storage {
	if a.storage == nil {
		db, err := storage.New(ctx, a.config.Postgres, a.log)
		if err != nil {
			a.log.Fatal("unable to initialize storage", zap.Error(err))
		}

		a.storage = db
		a.closer = append(a.closer, db.Close)
	}

	return a.storage
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
