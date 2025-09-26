package main

import (
	"flag"
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/joho/godotenv"
	"github.com/ksusonic/niti/backend/internal/migrations"
	"github.com/ksusonic/niti/backend/pgk/config"
	"github.com/ksusonic/niti/backend/pgk/logger"
	"go.uber.org/zap"
)

func main() {
	command := flag.String("command", "up", "Migration command: up, version")
	flag.Parse()

	_ = godotenv.Overload()

	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Errorf("load config: %v", err))
	}

	log, err := logger.New(cfg.Logger)
	if err != nil {
		panic(fmt.Errorf("create logger: %v", err))
	}

	switch *command {
	case "up":
		err = migrations.MigrateUp(cfg.Postgres)
		if err != nil {
			log.Fatal("run migrations", zap.Error(err))
		}
		log.Info("Migrations applied successfully")
	case "drop":
		var response string
		fmt.Println("Are you sure you want to drop all migrations? This action cannot be undone. (y/n): ")
		_, err := fmt.Scanln(&response)
		if err != nil || response != "y" && response != "Y" {
			log.Info("Drop migrations cancelled.")
			return
		}

		err = migrations.MigrateDrop(cfg.Postgres)
		if err != nil {
			log.Fatal("run migrations", zap.Error(err))
		}
		log.Info("Migrations rolled back successfully")
	case "version":
		migrator, err := migrations.NewMigrator(cfg.Postgres)
		if err != nil {
			log.Fatal("create migrator", zap.Error(err))
		}
		defer func() { _ = migrator.Close() }()

		currentVersion, dirty, err := migrator.Version()
		if err != nil {
			log.Fatal("get version", zap.Error(err))
		}
		if dirty {
			fmt.Printf("current version: %d (dirty)\n", currentVersion)
		} else {
			fmt.Printf("current version: %d\n", currentVersion)
		}

	default:
		fmt.Printf("unknown command: %s. Available commands: up, version", *command)
	}
}
