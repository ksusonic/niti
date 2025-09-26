package auth

import (
	"fmt"

	"github.com/ksusonic/niti/backend/internal/models"
	"github.com/ksusonic/niti/backend/internal/utils"
	initdata "github.com/telegram-mini-apps/init-data-golang"
)

func (s *Service) ParseInitData(data string) (*models.User, error) {
	if data == "" {
		return nil, fmt.Errorf("empty initData")
	}

	if err := initdata.Validate(data, s.cfg.TelegramToken, s.cfg.TokenExpiresIn); err != nil {
		return nil, fmt.Errorf("validate token: %w", err)
	}

	parsed, err := initdata.Parse(data)
	if err != nil {
		return nil, fmt.Errorf("parse initData: %w", err)
	}

	return &models.User{
		TelegramID: parsed.User.ID,
		Username:   parsed.User.Username,
		FirstName:  parsed.User.FirstName,
		LastName:   utils.NilIfEmpty(parsed.User.LastName),
		AvatarURL:  utils.NilIfEmpty(parsed.User.PhotoURL),
		IsDJ:       false,
	}, nil
}
