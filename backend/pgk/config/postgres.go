package config

type PostgresConfig struct {
	DSN string `env:"POSTGRES_DSN"`
}
