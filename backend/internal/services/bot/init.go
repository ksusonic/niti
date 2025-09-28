package bot

import (
	"fmt"
	"time"

	"github.com/ksusonic/niti/backend/pkg/config"
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

	return b, nil
}
