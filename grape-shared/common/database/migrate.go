package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/spf13/viper"
)

const defaultMigrationsDir = "./migrations"

// RunMigrations runs database migrations using Goose.
// For the auth service (grape_auth), set MIGRATIONS_DIR=./migrations/auth so only auth tables are created.
func RunMigrations() error {
	dir := os.Getenv("MIGRATIONS_DIR")
	if dir == "" {
		dir = viper.GetString("migrations.dir")
	}
	if dir == "" {
		dir = defaultMigrationsDir
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.dbname"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set dialect: %w", err)
	}

	if err := goose.Up(db, dir); err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	return nil
}
