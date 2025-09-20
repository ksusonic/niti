package repository

import "context"

func (r *RefreshTokenRepository) DeleteExpired(ctx context.Context) (int64, error) {
	res, err := r.Exec(
		ctx,
		`DELETE FROM refresh_tokens
		WHERE expires_at < now()`,
	)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected(), nil
}
