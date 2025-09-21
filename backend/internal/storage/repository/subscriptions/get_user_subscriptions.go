package subscriptions

import "context"

const eventsLimit = 30

func (r *Repository) GetUserSubscriptions(ctx context.Context, userID int64) ([]int, error) {
	rows, err := r.Query(
		ctx,
		`SELECT
			event_id
		FROM
			subscriptions
		WHERE
			user_id = $1
		ORDER BY
			create_time DESC
		LIMIT
			$2`,
		userID,
		eventsLimit,
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
