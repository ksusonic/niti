package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ksusonic/niti/backend/pkg/config"
	"go.uber.org/zap"
)

type Storage struct {
	*pgxpool.Pool
}

func New(ctx context.Context, cfg config.PostgresConfig, logger *zap.Logger) (*Storage, error) {
	if cfg.DSN == "" {
		return nil, fmt.Errorf("PostgreSQL DSN is not configured")
	}

	logger.Debug("connecting to postgres...")

	pool, err := pgxpool.New(ctx, cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("connect to postgres: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	logger.Info("postgres health check ok")

	return &Storage{
		Pool: pool,
	}, nil
}

func (s *Storage) Close(_ context.Context) {
	s.Pool.Close()
}
