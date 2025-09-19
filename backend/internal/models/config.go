package models

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken     string
	TelegramExpiresIn time.Duration

	ServerPort string

	AccessSecret  []byte
	RefreshSecret []byte
	AccessTTL     time.Duration
	RefreshTTL    time.Duration
}

func LoadConfig() (*Config, error) {
	// Load .env file if present
	_ = godotenv.Load()

	config := &Config{}

	// Load Telegram token (required)
	config.TelegramToken = os.Getenv("TELEGRAM_BOT_TOKEN")
	if config.TelegramToken == "" {
		return nil, fmt.Errorf("TELEGRAM_BOT_TOKEN environment variable is required")
	}

	// Load expiration duration (default: 24 hours)
	expiresInStr := os.Getenv("TELEGRAM_EXPIRES_IN_SECONDS")
	if expiresInStr == "" {
		config.TelegramExpiresIn = 24 * time.Hour
	} else {
		seconds, err := strconv.Atoi(expiresInStr)
		if err != nil {
			return nil, fmt.Errorf("invalid TELEGRAM_EXPIRES_IN_SECONDS value: %w", err)
		}
		config.TelegramExpiresIn = time.Duration(seconds) * time.Second
	}

	// Load server port (default: 8080)
	config.ServerPort = os.Getenv("SERVER_PORT")
	if config.ServerPort == "" {
		config.ServerPort = "8080"
	}

	// Load AccessSecret (default: "super_secret_access")
	accessSecret := os.Getenv("ACCESS_SECRET")
	if accessSecret == "" {
		accessSecret = "super_secret_access"
	}
	config.AccessSecret = []byte(accessSecret)

	// Load RefreshSecret (default: "super_secret_refresh")
	refreshSecret := os.Getenv("REFRESH_SECRET")
	if refreshSecret == "" {
		refreshSecret = "super_secret_refresh"
	}
	config.RefreshSecret = []byte(refreshSecret)

	// Load AccessTTL (default: 15 minutes)
	accessTTLStr := os.Getenv("ACCESS_TTL_MINUTES")
	if accessTTLStr == "" {
		config.AccessTTL = 15 * time.Minute
	} else {
		minutes, err := strconv.Atoi(accessTTLStr)
		if err != nil {
			return nil, fmt.Errorf("invalid ACCESS_TTL_MINUTES value: %w", err)
		}
		config.AccessTTL = time.Duration(minutes) * time.Minute
	}

	// Load RefreshTTL (default: 7 days)
	refreshTTLStr := os.Getenv("REFRESH_TTL_HOURS")
	if refreshTTLStr == "" {
		config.RefreshTTL = 7 * 24 * time.Hour
	} else {
		hours, err := strconv.Atoi(refreshTTLStr)
		if err != nil {
			return nil, fmt.Errorf("invalid REFRESH_TTL_HOURS value: %w", err)
		}
		config.RefreshTTL = time.Duration(hours) * time.Hour
	}

	return config, nil
}
