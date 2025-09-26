package bot

import (
	"fmt"
	"time"

	"github.com/ksusonic/niti/backend/pgk/config"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v4"
)

func initBot(
	cfg config.TelegramConfig,
	log *zap.Logger,
) (*tele.Bot, error) {
	pref := tele.Settings{
		Token:   cfg.Token,
		Poller:  &tele.LongPoller{Timeout: 10 * time.Second},
		Verbose: cfg.Verbose,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		return nil, fmt.Errorf("tele.NewBot: %w", err)
	}

	log.Info("setting menu button")

	err = b.SetMenuButton(nil, &tele.MenuButton{
		Type: tele.MenuButtonWebApp,
		Text: appTitle,
		WebApp: &tele.WebApp{
			URL: cfg.WebAppUrl,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("bot set menu button: %w", err)
	}

	return b, nil
}
