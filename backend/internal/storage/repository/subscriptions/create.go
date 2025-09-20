package subscriptions

import "context"

func (r *Repository) CreateSubscription(ctx context.Context, userID int64, eventID int) error {
	_, err := r.Exec(
		ctx,
		`INSERT INTO subscriptions (
			user_id,
			event_id
		)
		VALUES ($1, $2)
		ON CONFLICT (user_id, event_id) DO NOTHING`,
		userID,
		eventID,
	)

	return err
}
