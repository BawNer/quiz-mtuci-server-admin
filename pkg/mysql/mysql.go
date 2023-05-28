package mysql

import (
	"fmt"
	"quiz-mtuci-server/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQL -.
type MySQL struct {
	DB      *gorm.DB
	stopped bool
}

// New -.
func New(cfg config.MySQL) (*MySQL, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.DbName),
		SkipInitializeWithVersion: false,
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

	return &MySQL{
		DB: db,
	}, nil
}

// Close -.
func (m *MySQL) Close() {
	m.stopped = true
	sqlDB, err := m.DB.DB()

	if err == nil && sqlDB != nil {
		sqlDB.Close()
	}
}
