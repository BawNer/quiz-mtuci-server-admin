package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"quiz-mtuci-server/config"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const migrationsPath = "/migrations"

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("error %s", err)
	}

	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.DbName, cfg.Postgres.SSLMode)

	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatalf("error %s", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		log.Fatalf("error %s", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsPath,
		"postgres", driver)
	if err != nil {
		log.Fatalf("error %s", err)
	}
	defer m.Close()

	err = m.Up()
	if err == nil {
		log.Printf("Migrate: up success")
		return
	}
	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("Migrate: no change")
		return
	}
	{
		log.Fatalf("Migrate: up error: %s", err)
	}
}
