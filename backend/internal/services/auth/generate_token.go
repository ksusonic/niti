package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/ksusonic/niti/backend/internal/models"
)

func (s *Service) GenerateToken(ctx context.Context, telegramUserID int64) (models.JWTAuth, error) {
	now := time.Now()

	// access
	ac := accessClaims{
		UserID: telegramUserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(s.cfg.AccessTTL)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, ac)
	accessToken, err := at.SignedString(s.cfg.AccessSecret)
	if err != nil {
		return models.JWTAuth{}, fmt.Errorf("sign access token: %w", err)
	}

	// refresh
	jti := uuid.New()
	refreshExpiresAt := now.Add(s.cfg.RefreshTTL)
	rc := refreshClaims{
		UserID: telegramUserID,
		JTI:    jti.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExpiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rc)
	refreshToken, err := rt.SignedString(s.cfg.RefreshSecret)
	if err != nil {
		return models.JWTAuth{}, fmt.Errorf("sign refresh token: %w", err)
	}

	err = s.refreshTokenRepo.Insert(ctx, models.RefreshToken{
		JTI:       jti,
		UserID:    telegramUserID,
		ExpiresAt: refreshExpiresAt,
	})
	if err != nil {
		return models.JWTAuth{}, fmt.Errorf("insert refresh token: %w", err)
	}

	return models.JWTAuth{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		JTI:          jti,
	}, nil
}
