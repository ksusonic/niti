package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/ksusonic/niti/backend/pgk/config"
	"go.uber.org/zap"
)

type Storage struct {
	conn   *pgx.Conn
	logger *zap.Logger
}

func New(ctx context.Context, cfg config.PostgresConfig, logger *zap.Logger) (*Storage, error) {
	if cfg.DSN == "" {
		return nil, fmt.Errorf("PostgreSQL DSN is not configured")
	}

	logger.Debug("connecting to postgres")

	conn, err := pgx.Connect(ctx, cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("connect to postgres: %w", err)
	}

	logger.Info("successfully connected to postgres")

	if err := conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ping postgres: %w", err)
	}

	return &Storage{
		conn:   conn,
		logger: logger,
	}, nil
}

func (s *Storage) Close(ctx context.Context) error {
	return s.conn.Close(ctx)
}

func (s *Storage) GetConn() *pgx.Conn {
	return s.conn
}
