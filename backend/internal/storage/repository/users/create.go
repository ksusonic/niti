package users

import (
	"context"
	"fmt"

	"github.com/ksusonic/niti/backend/internal/models"
)

const createQuery = `
		INSERT INTO users (
			telegram_id,
			username,
			first_name,
			last_name,
			avatar_url,
			is_dj
		)
		VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (telegram_id) DO UPDATE SET
			username = EXCLUDED.username,
			first_name = EXCLUDED.first_name,
			last_name = EXCLUDED.last_name,
			avatar_url = EXCLUDED.avatar_url,
			updated_at = now()
		RETURNING
			telegram_id,
			username,
			first_name,
			last_name,
			avatar_url,
			is_dj`

func (r *Repository) Create(ctx context.Context, in *models.User) (*models.User, error) {
	rows, err := r.Query(
		ctx,
		createQuery,
		in.TelegramID,
		in.Username,
		in.FirstName,
		in.LastName,
		in.AvatarURL,
		in.IsDJ,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, fmt.Errorf("no rows returned")
	}

	var out models.User
	return &out, rows.Scan(
		&out.TelegramID,
		&out.Username,
		&out.FirstName,
		&out.LastName,
		&out.AvatarURL,
		&out.IsDJ,
	)
}
