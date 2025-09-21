package subscriptions

import "context"

func (r *Repository) DeleteSubscription(ctx context.Context, userID int, eventID int) error {
	_, err := r.Exec(
		ctx,
		`DELETE FROM subscriptions
		WHERE user_id = $1 AND event_id = $2`,
		userID,
		eventID,
	)

	return err
}
