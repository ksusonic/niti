package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/ksusonic/niti/backend/internal/storage/repository/base"
)

type Event struct {
	ID          int
	Title       string
	Description *string
	Location    *string
	VideoURL    *string
	StartsAt    string // Use time.Time if you have it imported
	CreatedBy   int64
}

type Repository struct {
	*base.Repository
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{Repository: base.NewBaseRepository(pool)}
}

// Example CRUD methods

func (r *Repository) CreateEvent(ctx context.Context, event *Event) (int, error) {
	var id int
	err := r.Pool.QueryRow(
		ctx,
		`INSERT INTO events (title, description, location, video_url, starts_at, created_by)
		 VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		event.Title, event.Description, event.Location, event.VideoURL, event.StartsAt, event.CreatedBy,
	).Scan(&id)
	return id, err
}

func (r *Repository) GetEventByID(ctx context.Context, id int) (*Event, error) {
	row := r.Pool.QueryRow(
		ctx,
		`SELECT id, title, description, location, video_url, starts_at, created_by FROM events WHERE id = $1`,
		id,
	)
	var event Event
	err := row.Scan(
		&event.ID,
		&event.Title,
		&event.Description,
		&event.Location,
		&event.VideoURL,
		&event.StartsAt,
		&event.CreatedBy,
	)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *Repository) ListEvents(ctx context.Context) ([]*Event, error) {
	rows, err := r.Pool.Query(
		ctx,
		`SELECT id, title, description, location, video_url, starts_at, created_by FROM events`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*Event
	for rows.Next() {
		var event Event
		err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.Location,
			&event.VideoURL,
			&event.StartsAt,
			&event.CreatedBy,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	return events, nil
}

func (r *Repository) DeleteEvent(ctx context.Context, id int) error {
	_, err := r.Pool.Exec(ctx, `DELETE FROM events WHERE id = $1`, id)
	return err
}
