package subscriptions

import (
	"context"

	"github.com/ksusonic/niti/backend/internal/models"
)

// GetAllSubscriptions returns all subscriptions ordered by creation time with pagination support
// If userID is provided (not nil), filters subscriptions for that specific user
func (r *Repository) GetAllSubscriptions(ctx context.Context, userID *int64, offset, limit int) ([]models.Subscription, error) {
	if limit <= 0 {
		limit = 50 // default limit
	}
	if offset < 0 {
		offset = 0
	}

	var query string
	var args []interface{}

	if userID != nil {
		query = `SELECT user_id, event_id, created_at
			FROM subscriptions
			WHERE user_id = $1
			ORDER BY created_at DESC, event_id
			LIMIT $2 OFFSET $3`
		args = []interface{}{*userID, limit, offset}
	} else {
		query = `SELECT user_id, event_id, created_at
			FROM subscriptions
			ORDER BY created_at DESC, user_id, event_id
			LIMIT $1 OFFSET $2`
		args = []interface{}{limit, offset}
	}

	rows, err := r.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscriptions []models.Subscription
	for rows.Next() {
		var subscription models.Subscription
		if err := rows.Scan(&subscription.UserID, &subscription.EventID, &subscription.CreatedAt); err != nil {
			return nil, err
		}
		subscriptions = append(subscriptions, subscription)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return subscriptions, nil
}
