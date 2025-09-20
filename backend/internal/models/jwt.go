package models

import (
	"time"

	"github.com/google/uuid"
)

type JWTAuth struct {
	AccessToken  string
	RefreshToken string
	JTI          uuid.UUID
}

type RefreshToken struct {
	JTI       uuid.UUID
	UserID    int64
	ExpiresAt time.Time
	Revoked   bool
	CreatedAt time.Time
}
