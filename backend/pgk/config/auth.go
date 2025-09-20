package config

import "time"

type AuthConfig struct {
	// JWT
	AccessSecret  string        `env:"ACCESS_SECRET" envDefault:"super_secret_access"`
	RefreshSecret string        `env:"REFRESH_SECRET" envDefault:"super_secret_refresh"`
	AccessTTL     time.Duration `env:"ACCESS_TTL_MINUTES" envDefault:"15m"`
	RefreshTTL    time.Duration `env:"REFRESH_TTL_HOURS" envDefault:"168h"`

	// Telegram
	TelegramToken     string        `env:"TELEGRAM_BOT_TOKEN,required"`
	TelegramExpiresIn time.Duration `env:"TELEGRAM_EXPIRES_IN_SECONDS" envDefault:"86400s"`
}
