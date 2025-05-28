package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func Migrate(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not instantiate database driver: %w", err)
	}
	migrator, err := migrate.NewWithDatabaseInstance(
		"file:///app/db/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		return fmt.Errorf("could not instantiate migrator: %w", err)
	}

	err = migrator.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("could not run migrations: %w", err)
	}
	return nil
}
