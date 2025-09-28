package config

import "time"

type AuthConfig struct {
	// JWT
	AccessSecret  string        `env:"ACCESS_SECRET"`
	RefreshSecret string        `env:"REFRESH_SECRET"`
	AccessTTL     time.Duration `env:"ACCESS_TTL" envDefault:"15m"`
	RefreshTTL    time.Duration `env:"REFRESH_TTL" envDefault:"24h"`

	// Telegram
	TelegramToken  string        `env:"TELEGRAM_BOT_TOKEN"`
	TokenExpiresIn time.Duration `env:"TOKEN_EXPIRES_IN" envDefault:"86400s"`
}
