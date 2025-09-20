package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func (s *Service) ValidateAccessToken(tokenStr string) (int64, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &accessClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return s.cfg.AccessSecret, nil
	})
	if err != nil {
		return 0, fmt.Errorf("parse token: %w", err)
	}

	claims, ok := token.Claims.(*accessClaims)
	if !ok || !token.Valid {
		return 0, fmt.Errorf("invalid access token")
	}

	return claims.UserID, nil
}
