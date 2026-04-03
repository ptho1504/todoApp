package db

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	mymysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(db *sql.DB) error {
	driver, err := mymysql.WithInstance(db, &mymysql.Config{})
	if err != nil {
		return fmt.Errorf("migration driver error: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)
	if err != nil {
		return fmt.Errorf("migration init error: %w", err)
	}

	if err := m.Up(); err != nil && err.Error() != "no change" {
		return fmt.Errorf("migration failed: %w", err)
	}

	fmt.Println("✅ Migrations applied")
	return nil
}
