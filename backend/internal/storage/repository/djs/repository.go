package djs

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ksusonic/niti/backend/internal/storage/repository/base"
	"go.uber.org/zap"
)

type Repository struct {
	*base.Repository
	logger *zap.Logger
}

func New(pool *pgxpool.Pool, logger *zap.Logger) *Repository {
	return &Repository{
		Repository: base.NewBaseRepository(pool),
		logger:     logger,
	}
}
