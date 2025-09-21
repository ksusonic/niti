package events

import (
	"context"
	"encoding/json"

	"github.com/ksusonic/niti/backend/internal/models"
)

const (
	eventsLimit        = 30
	getUserEventsQuery = `
		SELECT
			e.id,
			e.title,
			e.description,
			e.location,
			e.cover_url,
			e.video_url,
			e.starts_at,
			e.created_by,
			COALESCE(participants.count, 0) as participants_count,
			COALESCE(dj_lineup.djs_json, '[]'::jsonb) as djs_json
		FROM
			subscriptions s LEFT JOIN events e
				on s.event_id = e.id
			LEFT JOIN (
				SELECT
					event_id,
					COUNT(*) as count
				FROM
					subscriptions
				GROUP BY
					event_id
			) participants
			ON e.id = participants.event_id
			LEFT JOIN (
				SELECT
					ed.event_id,
					jsonb_agg(
						jsonb_build_object(
							'stage_name', d.stage_name,
							'avatar_url', COALESCE(d.avatar_url, ''),
							'socials', CASE
								WHEN d.socials IS NULL THEN '[]'::jsonb
								ELSE d.socials
							END
						) ORDER BY ed.order_in_lineup
					) as djs_json
				FROM
					event_djs ed
					JOIN djs d ON ed.dj_id = d.id
				GROUP BY
					ed.event_id
			) dj_lineup
			ON e.id = dj_lineup.event_id
		WHERE
			user_id = $1
		ORDER BY
			e.starts_at DESC
		LIMIT
			$2`
)

func (r *Repository) GetUserEvents(
	ctx context.Context,
	userID int64,
) ([]models.EventEnriched, error) {
	rows, err := r.Query(
		ctx,
		getUserEventsQuery,
		userID,
		eventsLimit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	events := make([]models.EventEnriched, 0, len(rows.RawValues()))
	for rows.Next() {
		event := models.EventEnriched{
			IsSubscribed: true,
		}
		var djsJSON []byte

		if err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.Location,
			&event.CoverURL,
			&event.VideoURL,
			&event.StartsAt,
			&event.CreatedBy,
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
