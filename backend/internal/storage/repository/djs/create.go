package djs

import (
	"context"
	"encoding/json"

	"github.com/ksusonic/niti/backend/internal/models"
	"go.uber.org/zap"
)

func (r *Repository) Create(ctx context.Context, dj *models.DJ) (*models.DJ, error) {
	socialsJSON, err := json.Marshal(dj.Socials)
	if err != nil {
		return nil, err
	}

	var created models.DJ
	var socialsBytes []byte

	err = r.QueryRow(
		ctx,
		`INSERT INTO djs (
			user_id,
			stage_name,
			avatar_url,
			socials
		)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (user_id) DO UPDATE SET
			stage_name = EXCLUDED.stage_name,
			avatar_url = EXCLUDED.avatar_url,
			socials = EXCLUDED.socials
		RETURNING
			id,
			user_id,
			stage_name,
			avatar_url,
			socials
		`, dj.TelegramID, dj.StageName, dj.AvatarURL, string(socialsJSON)).
		Scan(
			&created.ID,
			&created.TelegramID,
			&created.StageName,
			&created.AvatarURL,
			&socialsBytes,
		)
	if err != nil {
		return nil, err
	}

	if len(socialsBytes) > 0 {
		if err := json.Unmarshal(socialsBytes, &created.Socials); err != nil {
			r.logger.Error("unmarshal dj socials", zap.Error(err))
			created.Socials = []models.Social{}
		}
	} else {
		created.Socials = []models.Social{}
	}

	return &created, nil
}
