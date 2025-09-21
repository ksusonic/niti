package base

import (
	"testing"

	"github.com/caarlos0/env/v11"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/ksusonic/niti/backend/pgk/config"
	"github.com/stretchr/testify/require"
)

func SetupTestPool(t *testing.T) *pgxpool.Pool {
	require.NoError(t, godotenv.Load("../../../../.env"))

	var cfg config.PostgresConfig
	require.NoError(t, env.Parse(&cfg))

	config, err := pgxpool.ParseConfig(cfg.DSN)
	require.NoError(t, err)

	// Disable prepared statement caching to avoid conflicts in tests
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeExec
	config.MaxConns = 3

	pool, err := pgxpool.NewWithConfig(t.Context(), config)
	require.NoError(t, err)

	t.Cleanup(func() {
		pool.Close()
	})

	return pool
}
