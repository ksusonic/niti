package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/ksusonic/niti/backend/internal/models"
)

func (s *Service) ValidateRefreshToken(ctx context.Context, refreshTokenStr string) (*models.RefreshToken, error) {
	token, err := jwt.ParseWithClaims(refreshTokenStr, &refreshClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return s.cfg.RefreshSecret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("parse token: %w", err)
	}

	claims, ok := token.Claims.(*refreshClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid refresh token")
	}

	jti, err := uuid.Parse(claims.JTI)
	if err != nil {
		return nil, fmt.Errorf("invalid jti: %w", err)
	}

	refreshToken, err := s.refreshTokenRepo.GetValid(ctx, jti)
	if err != nil {
		return nil, fmt.Errorf("refresh token not found: %w", err)
	}
	if refreshToken.ExpiresAt.Before(time.Now()) {
		return nil, fmt.Errorf("refresh token expired")
	}

	return refreshToken, nil
}
