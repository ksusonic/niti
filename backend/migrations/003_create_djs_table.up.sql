CREATE TABLE djs (
    id SERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(telegram_id),
    stage_name TEXT NOT NULL,
    avatar_url TEXT,
    socials JSONB
);

CREATE INDEX idx_djs_stage_name ON djs(stage_name);
