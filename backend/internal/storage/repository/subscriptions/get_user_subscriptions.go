package subscriptions

import (
	"context"
)

// GetUserSubscriptions returns all event IDs that the user is subscribed to
func (r *Repository) GetUserSubscriptions(ctx context.Context, userID int64) ([]int, error) {
	rows, err := r.Query(
		ctx,
		`SELECT event_id FROM subscriptions
		WHERE user_id = $1
		ORDER BY event_id`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscribedEventIDs []int
	for rows.Next() {
		var eventID int
		if err := rows.Scan(&eventID); err != nil {
			return nil, err
		}
		subscribedEventIDs = append(subscribedEventIDs, eventID)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return subscribedEventIDs, nil
}
