package config

import "time"

type AuthConfig struct {
	// JWT
	AccessSecret  string        `env:"ACCESS_SECRET"`
	RefreshSecret string        `env:"REFRESH_SECRET"`
	AccessTTL     time.Duration `env:"ACCESS_TTL_MINUTES" envDefault:"15m"`
	RefreshTTL    time.Duration `env:"REFRESH_TTL_HOURS" envDefault:"168h"`

	// Telegram
	TelegramToken  string        `env:"TELEGRAM_BOT_TOKEN"`
	TokenExpiresIn time.Duration `env:"TOKEN_EXPIRES_IN" envDefault:"86400s"`
}
