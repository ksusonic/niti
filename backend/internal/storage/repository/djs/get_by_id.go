package djs

import (
	"context"
	"encoding/json"

	"github.com/jackc/pgx/v5"
	"github.com/ksusonic/niti/backend/internal/models"
	"go.uber.org/zap"
)

func (r *Repository) GetByID(ctx context.Context, telegramUserID int) (*models.DJ, error) {
	var dj models.DJ
	var socialsBytes []byte

	err := r.QueryRow(ctx, `
		SELECT
			id,
			user_id,
			stage_name,
			avatar_url,
			socials
		FROM
			djs
		WHERE
			user_id = $1`,
		telegramUserID,
	).Scan(
		&dj.ID,
		&dj.TelegramID,
		&dj.StageName,
		&dj.AvatarURL,
		&socialsBytes,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrNotFound
		}

		return nil, err
	}

	if len(socialsBytes) > 0 {
		if err := json.Unmarshal(socialsBytes, &dj.Socials); err != nil {
			r.logger.Error("unmarshal socials", zap.Error(err), zap.Int64("user_id", int64(telegramUserID)))
		}
	}

	return &dj, nil
}
