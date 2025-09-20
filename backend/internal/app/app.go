package app

import (
	"context"
	"fmt"

	"github.com/ksusonic/niti/backend/internal/auth"
	"github.com/ksusonic/niti/backend/internal/storage"
	"github.com/ksusonic/niti/backend/pgk/config"
	"github.com/ksusonic/niti/backend/pgk/logger"
	"go.uber.org/zap"
)

type App struct {
	config *config.Config
	log    *zap.Logger

	// lazy-init
	authService *auth.Service
	storage     *storage.Storage

	closer []func(context.Context) error
}

func New() *App {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Errorf("load config: %v", err))
	}

	log := logger.New(cfg.Logger)
	defer func() { _ = log.Sync() }()

	return &App{
		config: cfg,
		log:    log,
	}
}

func (a *App) Close(ctx context.Context) {
	for _, closeFunc := range a.closer {
		err := closeFunc(ctx)
		if err != nil {
			a.log.Error("close", zap.Error(err))
		}
	}
}

func (a *App) Config() *config.Config {
	return a.config
}

func (a *App) Logger() *zap.Logger {
	return a.log
}

func (a *App) Storage(ctx context.Context) *storage.Storage {
	if a.storage == nil {
		db, err := storage.New(ctx, a.config.Postgres, a.log)
		if err != nil {
			a.log.Fatal("unable to initialize storage", zap.Error(err))
		}

		a.storage = db
		a.closer = append(a.closer, db.Close)
	}

	return a.storage
}

func (a *App) AuthService() *auth.Service {
	if a.authService == nil {
		a.authService = auth.NewService(a.config.AuthConfig, a.log)
	}

	return a.authService
}
