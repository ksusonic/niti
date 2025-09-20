CREATE TABLE refresh_tokens (
    jti         UUID PRIMARY KEY,              -- уникальный ID токена (JWT ID)
    user_id     BIGINT NOT NULL,               -- ID пользователя
    expires_at  TIMESTAMPTZ NOT NULL,          -- срок жизни токена
    revoked     BOOLEAN NOT NULL DEFAULT false,-- отозван ли
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_refresh_tokens_user_id ON refresh_tokens(user_id);
CREATE INDEX idx_refresh_tokens_expires_at ON refresh_tokens(expires_at);
