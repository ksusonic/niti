package subscriptions

import (
	"context"
)

// GetSubscriptionsByEventIDs returns a slice of event IDs that the user is subscribed to
// from the provided list of event IDs
func (r *Repository) GetSubscriptionsByEventIDs(ctx context.Context, userID int, eventIDs []int) ([]int, error) {
	if len(eventIDs) == 0 {
		return []int{}, nil
	}

	rows, err := r.Query(
		ctx,
		`SELECT event_id FROM subscriptions
		WHERE user_id = $1 AND event_id = ANY($2)
		ORDER BY event_id`,
		userID,
		eventIDs,
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
