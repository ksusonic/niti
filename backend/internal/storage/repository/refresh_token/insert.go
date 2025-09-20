package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
)

func (r *RefreshTokenRepository) Insert(
	ctx context.Context,
	jti uuid.UUID,
	userID int64,
	expiresAt time.Time,
) error {
	_, err := r.Exec(
		ctx,
		`INSERT INTO refresh_tokens
			(jti, user_id, expires_at)
		VALUES ($1, $2, $3)`,
		jti, userID, expiresAt,
	)

	return err
}
