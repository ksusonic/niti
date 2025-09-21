package users

import "context"

func (r *Repository) Delete(ctx context.Context, telegramID int64) error {
	_, err := r.Exec(ctx, `DELETE FROM users WHERE telegram_id = $1`, telegramID)
	return err
}
