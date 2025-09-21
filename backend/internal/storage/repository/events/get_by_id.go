package events

import (
	"context"

	"github.com/ksusonic/niti/backend/internal/models"
)

func (r *Repository) GetEventByID(ctx context.Context, id int) (*models.Event, error) {
	row := r.QueryRow(
		ctx,
		`SELECT
			id,
			title,
			description,
			location,
			video_url,
			starts_at,
			created_by
		FROM
			events
		WHERE
			id = $1`,
		id,
	)
	var event models.Event
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
