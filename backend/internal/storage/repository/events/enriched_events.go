package events

import (
	"context"
	"encoding/json"

	"github.com/ksusonic/niti/backend/internal/models"
)

func (r *Repository) enrichedEvents(
	ctx context.Context,
	sql string,
	args ...any,
) ([]models.EventEnriched, error) {
	rows, err := r.Query(
		ctx,
		sql,
		args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := make([]models.EventEnriched, 0, len(rows.RawValues()))
	for rows.Next() {
		var (
			event   models.EventEnriched
			djsJSON []byte
		)

		if err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.Location,
			&event.CoverURL,
			&event.VideoURL,
			&event.StartsAt,
			&event.CreatedBy,
			&event.IsSubscribed,
			&event.ParticipantsCount,
			&djsJSON,
		); err != nil {
			return nil, err
		}

		event.DJs = make([]models.DJ, 0)
		if len(djsJSON) > 0 && string(djsJSON) != "null" {
			if err := json.Unmarshal(djsJSON, &event.DJs); err != nil {
				event.DJs = make([]models.DJ, 0)
			}
		}

		events = append(events, event)
	}

	return events, rows.Err()
}
