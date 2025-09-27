package bot

import (
	"context"
	"fmt"

	tele "gopkg.in/telebot.v4"
)

func (s *Service) Manage(ctx context.Context) error {
	s.log.Info("starting bot manage")

	err := s.bot.SetMenuButton(nil, &tele.MenuButton{
		Type: tele.MenuButtonWebApp,
		Text: appTitle,
		WebApp: &tele.WebApp{
			URL: s.webAppUrl,
		},
	})
	if err != nil {
		return fmt.Errorf("bot set menu button: %w", err)
	}

	return nil
}
