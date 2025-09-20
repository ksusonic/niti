package repository

import (
	"context"

	"github.com/ksusonic/niti/backend/internal/models"
)

func (r *Repository) Insert(
	ctx context.Context,
	token models.RefreshToken,
) error {
	_, err := r.Exec(
		ctx,
		`INSERT INTO refresh_tokens
			(jti, user_id, expires_at)
		VALUES ($1, $2, $3)`,
		token.JTI, token.UserID, token.ExpiresAt,
	)

	return err
}
