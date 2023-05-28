// Package postgres implements postgres connection.
package postgres

import (
	"fmt"
	"quiz-mtuci-server/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Postgres -.
type Postgres struct {
	DB      *gorm.DB
	stopped bool
}

// New -.
func New(cfg config.Postgres) (*Postgres, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=%s", cfg.User, cfg.Password, cfg.DbName, cfg.Host, cfg.Port, cfg.SSLMode),
		PreferSimpleProtocol: false, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, err
	}

	if cfg.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	}

	if cfg.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	}

	return &Postgres{
		DB: db,
	}, nil
}

// Close -.
func (p *Postgres) Close() {
	p.stopped = true

	sqlDB, err := p.DB.DB()

	if err == nil && sqlDB != nil {
		sqlDB.Close()
	}
}
