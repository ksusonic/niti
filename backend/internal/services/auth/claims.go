package auth

import "github.com/golang-jwt/jwt/v5"

type accessClaims struct {
	TgUserID int64 `json:"sub"`
	jwt.RegisteredClaims
}

type refreshClaims struct {
	TgUserID int64  `json:"sub"`
	JTI      string `json:"jti"`
	jwt.RegisteredClaims
}
