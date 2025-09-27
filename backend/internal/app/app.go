package app

import (
	"context"
	"fmt"

	"github.com/ksusonic/niti/backend/internal/storage"
	"github.com/ksusonic/niti/backend/pgk/config"
	"github.com/ksusonic/niti/backend/pgk/logger"
	"go.uber.org/zap"
)

type App struct {
	Config *config.Config
	Log    *zap.Logger

	// lazy-init
	storage      *storage.Storage
	repositories repositories
	services     services

	closer []func(context.Context)
}

func New() *App {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Errorf("load config: %v", err))
	}

	log, err := logger.New(cfg.Logger)
	if err != nil {
		panic(fmt.Errorf("create logger: %v", err))
	}

	return &App{
		Config: cfg,
		Log:    log,
	}
}

func (a *App) Close(ctx context.Context) {
	defer func() { _ = a.Log.Sync() }()

	for _, closeFunc := range a.closer {
		closeFunc(ctx)
	}
}
