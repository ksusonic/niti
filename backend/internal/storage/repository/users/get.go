package users

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/ksusonic/niti/backend/internal/models"
)

func (r *Repository) Get(ctx context.Context, telegramID int64) (*models.User, error) {
	row := r.QueryRow(
		ctx, `
		SELECT
			telegram_id,
			username,
			first_name,
			last_name,
			avatar_url,
			is_dj
		FROM
			users
		WHERE
			telegram_id = $1`,
		telegramID,
	)

	var user models.User
	err := row.Scan(
		&user.TelegramID,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.AvatarURL,
		&user.IsDJ,
	)
	if err == pgx.ErrNoRows {
		return nil, models.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
