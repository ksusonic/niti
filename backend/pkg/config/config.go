package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	// Webserver
	Webserver WebserverConfig

	// Telegram
	Telegram TelegramConfig

	// Logger
	Logger LoggerConfig

	// Postgres
	Postgres PostgresConfig

	// JWT
	AuthConfig AuthConfig
}

func LoadConfig() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return nil, fmt.Errorf("parse env config: %w", err)
	}

	return &cfg, nil
}
