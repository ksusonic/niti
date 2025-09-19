package auth

import (
	"fmt"

	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func (s *Service) ParseInitData(data string) (*initdata.InitData, error) {
	if data == "" {
		return nil, fmt.Errorf("empty initData")
	}

	if err := initdata.Validate(data, s.token, s.expiresIn); err != nil {
		return nil, fmt.Errorf("validate token: %w", err)
	}

	parsed, err := initdata.Parse(data)
	if err != nil {
		return nil, fmt.Errorf("parse initData: %w", err)
	}

	return &parsed, nil
}
