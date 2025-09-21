package events

import (
	"context"

	"github.com/ksusonic/niti/backend/internal/models"
)

func (r *Repository) ListEvents(ctx context.Context) ([]models.Event, error) {
	rows, err := r.Query(
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
			events`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := make([]models.Event, 0, len(rows.RawValues()))
	for rows.Next() {
		var event models.Event
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
		events = append(events, event)
	}
	return events, nil
}
