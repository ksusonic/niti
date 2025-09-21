package subscriptions

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ksusonic/niti/backend/internal/storage/repository/base"
)

type Repository struct {
	*base.Repository
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{Repository: base.NewBaseRepository(pool)}
}
