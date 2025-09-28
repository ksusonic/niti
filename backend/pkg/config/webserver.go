package config

type WebserverConfig struct {
	Port           int      `env:"PORT" envDefault:"8080"`
	TrustedProxies []string `env:"TRUSTED_PROXIES" envDefault:"127.0.0.1"`
}
