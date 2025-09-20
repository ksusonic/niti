CREATE TABLE users (
    telegram_id BIGINT PRIMARY KEY,
    username TEXT,
    first_name TEXT,
    last_name TEXT,
    avatar_url TEXT,
    is_dj BOOLEAN DEFAULT false
);

CREATE INDEX idx_users_username ON users(username);
