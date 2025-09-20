package main

import (
	"flag"
	"log"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/ksusonic/niti/backend/internal/migrations"
	"github.com/ksusonic/niti/backend/pgk/config"
)

func main() {
	var (
		command = flag.String("command", "up", "Migration command: up, version")
		force   = flag.Bool("force", false, "Force load config even if .env file is missing")
	)
	flag.Parse()

	cfg, err := config.LoadConfig()
	if err != nil && !*force {
		log.Fatalf("load config: %v", err)
	}

	switch *command {
	case "up":
		err = migrations.MigrateUp(cfg.Postgres)
		if err != nil {
			log.Fatalf("run migrations: %v", err)
		}
		log.Println("Migrations applied successfully")

	case "version":
		migrator, err := migrations.NewMigrator(cfg.Postgres)
		if err != nil {
			log.Fatalf("create migrator: %v", err)
		}
		defer func() { _ = migrator.Close() }()

		currentVersion, dirty, err := migrator.Version()
		if err != nil {
			log.Fatalf("get version: %v", err)
		}
		if dirty {
			log.Printf("current version: %d (dirty)\n", currentVersion)
		} else {
			log.Printf("current version: %d\n", currentVersion)
		}

	default:
		log.Fatalf("unknown command: %s. Available commands: up, version", *command)
	}
}
