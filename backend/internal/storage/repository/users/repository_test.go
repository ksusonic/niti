//go:build integration

package users_test

import (
	"context"
	"testing"

	"github.com/caarlos0/env/v11"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/ksusonic/niti/backend/internal/models"
	"github.com/ksusonic/niti/backend/internal/storage/repository/users"
	"github.com/ksusonic/niti/backend/internal/utils"
	"github.com/ksusonic/niti/backend/pgk/config"
	"github.com/stretchr/testify/require"
)

func setupTestPool(t *testing.T) *pgxpool.Pool {
	require.NoError(t, godotenv.Load("../../../../.env"))

	var cfg config.PostgresConfig
	require.NoError(t, env.Parse(&cfg))

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

func TestUsersRepository_CRUD(t *testing.T) {
	pool := setupTestPool(t)
	repo := users.New(pool)
	ctx := context.Background()

	telegramID := int64(uuid.New().ID())
	username := "testuser"
	firstName := "Test"
	lastName := "User"
	avatarURL := "https://example.com/avatar.png"
	isDJ := true

	// CREATE
	user := &models.User{
		TelegramID: telegramID,
		Username:   &username,
		FirstName:  &firstName,
		LastName:   &lastName,
		AvatarURL:  &avatarURL,
		IsDJ:       isDJ,
	}
	created, err := repo.Create(ctx, user)
	require.NoError(t, err)
	require.NotNil(t, created)
	require.Equal(t, telegramID, created.TelegramID)
	require.Equal(t, username, utils.Deref(created.Username))
	require.Equal(t, firstName, utils.Deref(created.FirstName))
	require.Equal(t, lastName, utils.Deref(created.LastName))
	require.Equal(t, avatarURL, utils.Deref(created.AvatarURL))
	require.Equal(t, isDJ, created.IsDJ)

	// GET
	got, err := repo.Get(ctx, telegramID)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, telegramID, got.TelegramID)
	require.Equal(t, username, utils.Deref(got.Username))
	require.Equal(t, firstName, utils.Deref(got.FirstName))
	require.Equal(t, lastName, utils.Deref(got.LastName))
	require.Equal(t, avatarURL, utils.Deref(got.AvatarURL))
	require.Equal(t, isDJ, got.IsDJ)

	// UPDATE (change username)
	newUsername := "updateduser"
	user.Username = &newUsername
	updated, err := repo.Create(ctx, user)
	require.NoError(t, err)
	require.NotNil(t, updated)
	require.Equal(t, newUsername, utils.Deref(updated.Username))

	// GET after update
	gotUpdated, err := repo.Get(ctx, telegramID)
	require.NoError(t, err)
	require.NotNil(t, gotUpdated)
	require.Equal(t, newUsername, utils.Deref(gotUpdated.Username))

	// DELETE
	err = repo.Delete(ctx, telegramID)
	require.NoError(t, err)

	// GET after delete (should be nil)
	gotDeleted, err := repo.Get(ctx, telegramID)
	require.NoError(t, err)
	require.Nil(t, gotDeleted)
}
