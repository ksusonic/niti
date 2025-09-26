package config

type TelegramConfig struct {
	Verbose   bool   `env:"TELEGRAM_BOT_VERBOSE"`
	Token     string `env:"TELEGRAM_BOT_TOKEN"`
	WebAppUrl string `env:"TELEGRAM_BOT_WEBAPP_URL" envDefault:"https://example.com"`
}
