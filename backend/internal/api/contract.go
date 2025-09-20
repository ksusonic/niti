//go:generate go tool mockgen -destination=./mocks/mock_auth.go -package=mocks -source=contract.go

package api

import (
	"context"

	"github.com/ksusonic/niti/backend/internal/models"
)

type auth interface {
	ParseInitData(initData string) (*models.User, error)
	GenerateTokens(ctx context.Context, userID int64) (models.JWTokens, error)
	ValidateRefreshToken(ctx context.Context, refreshTokenStr string) (*models.RefreshToken, error)
	RollTokens(ctx context.Context, refresh *models.RefreshToken) (*models.JWTokens, error)
}

type usersRepo interface {
	Get(ctx context.Context, telegramID int64) (*models.User, error)
	Create(ctx context.Context, in *models.User) (*models.User, error)
}

type subscriptionsRepo interface {
	CreateSubscription(ctx context.Context, userID int64, eventID int) error
}
