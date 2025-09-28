package bot

import (
	"context"

	"github.com/ksusonic/niti/backend/pkg/config"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v4"
)

type Service struct {
	bot *tele.Bot
	log *zap.Logger

	webAppUrl string
}

func NewService(
	cfg config.TelegramConfig,
	log *zap.Logger,
) *Service {
	if cfg.Token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN required")
	}

	bot, err := initBot(cfg, log)
	if err != nil {
		log.Fatal("init bot", zap.Error(err))
	}

	s := &Service{
		bot: bot,
		log: log,

		webAppUrl: cfg.WebAppUrl,
	}

	s.bot.Handle("/start", s.startCommand)

	return s
}

func (s *Service) Start(ctx context.Context) {
	go func() {
		s.log.Info("starting bot")
		s.bot.Start()
	}()

	<-ctx.Done()
	s.log.Info("stopping bot...")

	s.bot.Stop()
}
