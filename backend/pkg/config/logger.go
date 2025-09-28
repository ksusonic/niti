package config

type LoggerConfig struct {
	LogLevel  string `env:"LOG_LEVEL" envDefault:"debug"`
	LogFormat string `env:"LOG_FORMAT" envDefault:"simple"`
}
