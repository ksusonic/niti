package subscriptions

import (
	"context"
)

// GetSubscriptionsCount returns the total number of subscriptions
// If userID is provided (not nil), returns count for that specific user
func (r *Repository) GetSubscriptionsCount(ctx context.Context, userID *int64) (int64, error) {
	var count int64
	var query string
	var args []interface{}

	if userID != nil {
		query = `SELECT COUNT(*) FROM subscriptions WHERE user_id = $1`
		args = []interface{}{*userID}
	} else {
		query = `SELECT COUNT(*) FROM subscriptions`
		args = []interface{}{}
	}

	err := r.QueryRow(ctx, query, args...).Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}
