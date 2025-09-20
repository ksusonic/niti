package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ksusonic/niti/backend/internal/storage/repository/base"
)

type EventDJ struct {
	EventID       int
	DJID          int
	OrderInLineup int
}

type Repository struct {
	*base.Repository
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{Repository: base.NewBaseRepository(pool)}
}

// Create inserts a new event_djs record
func (r *Repository) Create(ctx context.Context, eventDJ EventDJ) error {
	_, err := r.Pool.Exec(ctx,
		`INSERT INTO event_djs (event_id, dj_id, order_in_lineup) VALUES ($1, $2, $3)`,
		eventDJ.EventID, eventDJ.DJID, eventDJ.OrderInLineup,
	)
	return err
}

// GetByEventID returns all DJs for a given event, ordered by lineup
func (r *Repository) GetByEventID(ctx context.Context, eventID int) ([]EventDJ, error) {
	rows, err := r.Pool.Query(ctx,
		`SELECT event_id, dj_id, order_in_lineup FROM event_djs WHERE event_id = $1 ORDER BY order_in_lineup ASC`,
		eventID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []EventDJ
	for rows.Next() {
		var edj EventDJ
		if err := rows.Scan(&edj.EventID, &edj.DJID, &edj.OrderInLineup); err != nil {
			return nil, err
		}
		result = append(result, edj)
	}
	return result, rows.Err()
}

// Delete removes a DJ from an event lineup
func (r *Repository) Delete(ctx context.Context, eventID, djID int) error {
	_, err := r.Pool.Exec(ctx,
		`DELETE FROM event_djs WHERE event_id = $1 AND dj_id = $2`,
		eventID, djID,
	)
	return err
}

// UpdateOrder updates the order_in_lineup for a DJ in an event
func (r *Repository) UpdateOrder(ctx context.Context, eventID, djID, orderInLineup int) error {
	_, err := r.Pool.Exec(ctx,
		`UPDATE event_djs SET order_in_lineup = $1 WHERE event_id = $2 AND dj_id = $3`,
		orderInLineup, eventID, djID,
	)
	return err
}
