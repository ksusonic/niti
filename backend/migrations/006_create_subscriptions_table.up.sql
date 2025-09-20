CREATE TABLE subscriptions (
    user_id BIGINT REFERENCES users(telegram_id) ON DELETE CASCADE,
    event_id INT REFERENCES events(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, event_id)
);

CREATE INDEX idx_subscriptions_event_id ON subscriptions(event_id);
CREATE INDEX idx_subscriptions_user_id ON subscriptions(user_id);
