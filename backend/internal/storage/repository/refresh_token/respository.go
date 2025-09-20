package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ksusonic/niti/backend/internal/storage/repository/base"
)

type RefreshTokenRepository struct {
	*base.Repository
}

func NewRefreshTokenRepository(pool *pgxpool.Pool) *RefreshTokenRepository {
	return &RefreshTokenRepository{Repository: base.NewBaseRepository(pool)}
}
