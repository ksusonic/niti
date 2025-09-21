//go:build integration

package djs_test

import (
	"context"
	"testing"

	"github.com/ksusonic/niti/backend/internal/models"
	"github.com/ksusonic/niti/backend/internal/storage/repository/base"
	"github.com/ksusonic/niti/backend/internal/storage/repository/djs"
	"github.com/ksusonic/niti/backend/internal/storage/repository/users"
	"github.com/ksusonic/niti/backend/internal/utils"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestDJsRepository(t *testing.T) {
	pool := base.SetupTestPool(t)
	logger := zap.NewNop()
	repo := djs.New(pool, logger)
	ctx := context.Background()

	telegramID := int64(123456789)
	stageName := "Test DJ"
	avatarURL := "https://example.com/dj-avatar.png"
	socials := []models.Social{
		{
			Name: "Instagram",
			URL:  "https://instagram.com/testdj",
			Icon: utils.Ptr("instagram"),
		},
		{
			Name: "SoundCloud",
			URL:  "https://soundcloud.com/testdj",
		},
	}

	err := repo.WithRollback(ctx, func(ctx context.Context) {
		// CREATE
		dj := &models.DJ{
			TelegramID: nil,
			StageName:  stageName,
			AvatarURL:  &avatarURL,
			Socials:    socials,
		}
		created, err := repo.Create(ctx, dj)
		require.NoError(t, err)
		require.NotNil(t, created)
		require.NotZero(t, created.ID)
		require.Equal(t, stageName, created.StageName)
		require.Equal(t, avatarURL, *created.AvatarURL)
		require.Len(t, created.Socials, 2)
		require.Equal(t, "Instagram", created.Socials[0].Name)
		require.Equal(t, "https://instagram.com/testdj", created.Socials[0].URL)
		require.NotNil(t, created.Socials[0].Icon)
		require.Equal(t, "instagram", *created.Socials[0].Icon)

		userRepo := users.New(pool)
		_, err = userRepo.Create(ctx, &models.User{
			TelegramID: telegramID,
		})
		require.NoError(t, err)

		// UPDATE ON INSERT
		dj.TelegramID = &telegramID
		updated, err := repo.Create(ctx, dj)
		require.NoError(t, err)
		require.NotNil(t, updated)
		require.NotNil(t, updated.TelegramID)
		require.Equal(t, telegramID, *updated.TelegramID)

		// GET
		got, err := repo.GetByID(ctx, int(telegramID))
		require.NoError(t, err)
		require.NotNil(t, got)
		require.Equal(t, telegramID, *got.TelegramID)
		require.Equal(t, stageName, got.StageName)
		require.Equal(t, avatarURL, *got.AvatarURL)
		require.Len(t, got.Socials, 2)
		require.Equal(t, "SoundCloud", got.Socials[1].Name)
		require.Nil(t, created.Socials[1].Icon)
	})
	require.NoError(t, err)
}
