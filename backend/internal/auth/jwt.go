package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/ksusonic/niti/backend/internal/models"
)

func (s *Service) GenerateToken(telegramUserID int64) (models.JWTAuth, error) {
	now := time.Now()

	// access
	ac := AccessClaims{
		TelegramUserID: telegramUserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(s.accessTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, ac)
	accessToken, err := at.SignedString(s.accessSecret)
	if err != nil {
		return models.JWTAuth{}, fmt.Errorf("sign access token: %w", err)
	}

	// refresh
	jti := uuid.NewString()
	rc := RefreshClaims{
		TelegramUserID: telegramUserID,
		JTI:            jti,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(s.refreshTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rc)
	refreshToken, err := rt.SignedString(s.refreshSecret)
	if err != nil {
		return models.JWTAuth{}, fmt.Errorf("sign refresh token: %w", err)
	}

	return models.JWTAuth{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		JTI:          jti,
	}, nil
}
