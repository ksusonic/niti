CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT,
    cover_url TEXT,
    video_url TEXT,
    starts_at TIMESTAMP,
    created_by BIGINT REFERENCES users(telegram_id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_events_starts_at ON events(starts_at);
CREATE INDEX idx_events_location ON events(location);
