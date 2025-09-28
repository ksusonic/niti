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

	parseConfig, err := pgxpool.ParseConfig(cfg.DSN)
	require.NoError(t, err)

	// Disable prepared statement caching to avoid conflicts in tests
	parseConfig.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeExec
	parseConfig.MaxConns = 3

	pool, err := pgxpool.NewWithConfig(t.Context(), parseConfig)
	require.NoError(t, err)

	t.Cleanup(func() {
		pool.Close()
	})

	return pool
}
