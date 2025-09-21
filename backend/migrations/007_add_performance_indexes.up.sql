CREATE INDEX idx_events_starts_at_id_desc ON events(starts_at DESC, id);

CREATE INDEX idx_subscriptions_user_event_created ON subscriptions(user_id, event_id, created_at);

CREATE INDEX idx_event_djs_event_order_dj ON event_djs(event_id, order_in_lineup, dj_id);

CREATE INDEX idx_events_covering ON events(id, starts_at DESC)
INCLUDE (title, description, location, cover_url, video_url, created_by);

CREATE INDEX idx_subscriptions_event_user ON subscriptions(event_id, user_id);

CREATE INDEX idx_event_djs_dj_event ON event_djs(dj_id, event_id);
