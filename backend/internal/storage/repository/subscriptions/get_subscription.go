package subscriptions

import (
	"context"
)

func (r *Repository) GetSubscription(ctx context.Context, userID int64, eventID int) (bool, error) {
	var exists bool
	err := r.QueryRow(
		ctx,
		`SELECT EXISTS(
			SELECT 1 FROM subscriptions
			WHERE user_id = $1 AND event_id = $2
		)`,
		userID,
		eventID,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}
