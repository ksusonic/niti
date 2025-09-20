package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	// Webserver
	ServerPort     string   `env:"PORT" envDefault:"8080"`
	TrustedProxies []string `env:"TRUSTED_PROXIES" envDefault:"127.0.0.1"`

	// Logger
	Logger LoggerConfig

	// Postgres
	Postgres PostgresConfig

	// JWT
	AuthConfig AuthConfig
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("load env: %w", err)
	}

	cfg := &Config{}
	err = env.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("parse env config: %w", err)
	}

	return cfg, nil
}
