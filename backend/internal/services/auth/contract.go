//go:generate go tool mockgen -destination=./mocks/mock_auth.go -package=mocks -source=contract.go

package auth

import (
	"context"

	"github.com/google/uuid"
	"github.com/ksusonic/niti/backend/internal/models"
)

type refreshTokenRepo interface {
	GetValid(ctx context.Context, jti uuid.UUID) (*models.RefreshToken, error)
	Insert(ctx context.Context, token models.RefreshToken) error
	Revoke(ctx context.Context, jti uuid.UUID) error
}
