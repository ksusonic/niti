package events

import (
	"context"

	"github.com/ksusonic/niti/backend/internal/models"
)

const listEventsQuery = `
	SELECT
		e.id,
		e.title,
		e.description,
		e.location,
		e.cover_url,
		e.video_url,
		e.starts_at,
		e.created_by,
		COALESCE(subscriptions.subscribed, false) as is_subscribed,
		COALESCE(participants.count, 0) as participants_count,
		COALESCE(dj_lineup.djs_json, '[]'::jsonb) as djs_json
	FROM
		events e
		LEFT JOIN (
			SELECT
				event_id,
				COUNT(*) as count
			FROM
				subscriptions
			GROUP BY
				event_id
		) participants ON e.id = participants.event_id
		LEFT JOIN (
			SELECT
				event_id,
				true as subscribed
			FROM
				subscriptions
			WHERE
				user_id = $1
		) subscriptions ON e.id = subscriptions.event_id
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
		) dj_lineup ON e.id = dj_lineup.event_id
	ORDER BY
		starts_at DESC
	LIMIT
		$2
	OFFSET
		$3`

func (r *Repository) ListEvents(
	ctx context.Context,
	userID int64,
	limit, offset int,
) ([]models.EventEnriched, error) {
	return r.enrichedEvents(
		ctx,
		listEventsQuery,
		userID,
		limit,
		offset,
	)
}
