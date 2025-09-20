package models

import (
	"time"

	"github.com/google/uuid"
)

type JWTokens struct {
	AccessToken  string
	RefreshToken string
	JTI          uuid.UUID
	ExpiresIn    time.Duration
}

type RefreshToken struct {
	JTI       uuid.UUID
	UserID    int64
	ExpiresAt time.Time
	Revoked   bool
	CreatedAt time.Time
}
