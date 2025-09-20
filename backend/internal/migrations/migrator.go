package migrations

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/ksusonic/niti/backend/pgk/config"
)

// Migrator manages database migrations
type Migrator struct {
	migrate *migrate.Migrate
}

// NewMigrator creates a new migrator instance using PostgresConfig
func NewMigrator(cfg config.PostgresConfig) (*Migrator, error) {
	if cfg.DSN == "" {
		return nil, fmt.Errorf("database DSN is required")
	}

	migrationsPath := cfg.MigrationsPath
	if migrationsPath == "" {
		migrationsPath = "migrations"
	}

	m, err := migrate.New(migrationsPath, cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to create migrate instance: %w", err)
	}

	return &Migrator{migrate: m}, nil
}

// NewMigratorWithDB creates a new migrator instance using an existing database connection
func NewMigratorWithDB(db *sql.DB, cfg config.PostgresConfig) (*Migrator, error) {
	migrationsPath := cfg.MigrationsPath
	if migrationsPath == "" {
		migrationsPath = "migrations"
	}

	driver, err := pgx.WithInstance(db, &pgx.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to create database driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(migrationsPath, "pgx5", driver)
	if err != nil {
		return nil, fmt.Errorf("failed to create migrate instance: %w", err)
	}

	return &Migrator{migrate: m}, nil
}

// Up runs all up migrations
func (m *Migrator) Up() error {
	err := m.migrate.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}
	return err
}

// Version returns the current migration version and dirty state
func (m *Migrator) Version() (version uint, dirty bool, err error) {
	return m.migrate.Version()
}

// Close closes the migrator instance
func (m *Migrator) Close() error {
	sourceErr, dbErr := m.migrate.Close()
	if sourceErr != nil {
		return sourceErr
	}
	return dbErr
}

// MigrateUp is a convenience function to run all up migrations
func MigrateUp(cfg config.PostgresConfig) error {
	migrator, err := NewMigrator(cfg)
	if err != nil {
		return err
	}
	defer func() { _ = migrator.Close() }()

	return migrator.Up()
}
