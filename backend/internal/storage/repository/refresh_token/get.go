package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/ksusonic/niti/backend/internal/models"
)

func (r *Repository) GetValid(ctx context.Context, jti uuid.UUID) (*models.RefreshToken, error) {
	row := r.QueryRow(
		ctx,
		`SELECT jti, user_id, expires_at, revoked, created_at
		 FROM refresh_tokens
		 WHERE
			jti = $1 AND
			revoked = false AND
			expires_at > now()`,
		jti,
	)

	var token models.RefreshToken
	err := row.Scan(&token.JTI, &token.UserID, &token.ExpiresAt, &token.Revoked, &token.CreatedAt)
	if err == pgx.ErrNoRows {
		return nil, models.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return &token, nil
}
