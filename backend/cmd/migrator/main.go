package main

import (
	"flag"
	"fmt"
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
		log.Fatalf("Failed to load config: %v", err)
	}

	switch *command {
	case "up":
		err = migrations.MigrateUp(cfg.Postgres)
		if err != nil {
			log.Fatalf("Failed to run migrations: %v", err)
		}
		fmt.Println("Migrations applied successfully")

	case "version":
		migrator, err := migrations.NewMigrator(cfg.Postgres)
		if err != nil {
			log.Fatalf("Failed to create migrator: %v", err)
		}
		defer migrator.Close()

		currentVersion, dirty, err := migrator.Version()
		if err != nil {
			log.Fatalf("Failed to get version: %v", err)
		}
		if dirty {
			fmt.Printf("Current version: %d (dirty)\n", currentVersion)
		} else {
			fmt.Printf("Current version: %d\n", currentVersion)
		}

	default:
		log.Fatalf("Unknown command: %s. Available commands: up, version", *command)
	}
}
