package database

import (
	"fmt"

	"github.com/lukestarwalk/url-shorter/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DbName   string
	SSLMode  string
}

func NewGormDB(dc *DBConfig) (*gorm.DB, error) {
	DSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dc.User, dc.Password, dc.Host, dc.Port, dc.DbName, dc.SSLMode)

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&model.URL{})
	return db, nil
}
