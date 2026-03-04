package database

import (
	"github.com/lukestarwalk/url-shorter/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&model.URL{})
	return db, nil
}
