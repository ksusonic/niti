package auth

import "github.com/golang-jwt/jwt/v5"

type AccessClaims struct {
	TelegramUserID int64 `json:"sub"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	TelegramUserID int64  `json:"sub"`
	JTI            string `json:"jti"`
	jwt.RegisteredClaims
}
