package repository

import (
	"context"

	"github.com/google/uuid"
)

func (r *Repository) Revoke(ctx context.Context, jti uuid.UUID) error {
	_, err := r.Exec(
		ctx,
		`UPDATE
			refresh_tokens
		SET
			revoked = true
		WHERE
			jti = $1`,
		jti,
	)
	return err
}
