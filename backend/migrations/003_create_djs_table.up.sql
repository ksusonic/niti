CREATE TABLE djs (
    id SERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(telegram_id),
    stage_name TEXT NOT NULL,
    avatar_url TEXT,
    socials JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX idx_djs_user ON djs(user_id);
CREATE INDEX idx_djs_stage_name ON djs(stage_name);
