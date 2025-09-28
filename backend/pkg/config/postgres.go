package config

type PostgresConfig struct {
	DSN            string `env:"POSTGRES_DSN"`
	MigrationsPath string `env:"POSTGRES_MIGRATIONS_PATH" envDefault:"file://migrations"`
}
