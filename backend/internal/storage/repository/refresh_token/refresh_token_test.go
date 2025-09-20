//go:build integration

package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	repository "github.com/ksusonic/niti/backend/internal/storage/repository/refresh_token"
	"github.com/ksusonic/niti/backend/pgk/config"
	"github.com/stretchr/testify/require"
)

func setupTestPool(t *testing.T) *pgxpool.Pool {
	require.NoError(t, godotenv.Load("../../../../.env"))

	var cfg config.PostgresConfig
	require.NoError(t, env.Parse(&cfg))

	// Parse the DSN and configure to disable prepared statement caching
	config, err := pgxpool.ParseConfig(cfg.DSN)
	require.NoError(t, err)

	// Disable prepared statement caching to avoid conflicts in tests
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeExec

	pool, err := pgxpool.NewWithConfig(t.Context(), config)
	require.NoError(t, err)

	t.Cleanup(func() {
		pool.Close()
	})

	return pool
}

func TestRefreshTokenRepository_CRUD(t *testing.T) {
	pool := setupTestPool(t)

	repo := repository.NewRefreshTokenRepository(pool)
	ctx := context.Background()

	jti := uuid.New()
	userID := int64(42)
	expiresAt := time.Now().Add(1 * time.Hour)

	err := repo.WithRollback(ctx, func(ctx context.Context) {
		// INSERT
		err := repo.Insert(ctx, jti, userID, expiresAt)
		require.NoError(t, err)

		// SELECT valid
		token, err := repo.GetValid(ctx, jti)
		require.NoError(t, err)
		require.NotNil(t, token)
		require.Equal(t, jti, token.JTI)
		require.Equal(t, userID, token.UserID)
		require.False(t, token.Revoked)
		require.True(t, token.ExpiresAt.After(time.Now()))

		// UPDATE (revoke)
		err = repo.Revoke(ctx, jti)
		require.NoError(t, err)

		// SELECT after revoke (should be nil)
		token, err = repo.GetValid(ctx, jti)
		require.NoError(t, err)
		require.Nil(t, token)

		// Simulate expiration for DELETE
		_, err = repo.Exec(ctx, `UPDATE refresh_tokens SET expires_at = now() - interval '1 minute' WHERE jti = $1`, jti)
		require.NoError(t, err)

		// DELETE expired
		affected, err := repo.DeleteExpired(ctx)
		require.NoError(t, err)
		require.GreaterOrEqual(t, affected, int64(1))
	})
	require.NoError(t, err)
}
