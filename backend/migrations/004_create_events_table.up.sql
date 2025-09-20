CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    location TEXT,
    video_url TEXT,
    starts_at TIMESTAMP NOT NULL,
    created_by BIGINT REFERENCES users(telegram_id)
);

CREATE INDEX idx_events_starts_at ON events(starts_at);
CREATE INDEX idx_events_location ON events(location);
