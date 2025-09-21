package events

import (
	"context"

	"github.com/ksusonic/niti/backend/internal/models"
)

func (r *Repository) CreateEvent(ctx context.Context, event *models.Event) (int, error) {
	var id int
	err := r.QueryRow(
		ctx,
		`INSERT INTO events (title, description, location, video_url, starts_at, created_by)
		 VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		event.Title, event.Description, event.Location, event.VideoURL, event.StartsAt, event.CreatedBy,
	).Scan(&id)
	return id, err
}
