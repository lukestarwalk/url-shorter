package repository

import (
	"github.com/lukestarwalk/url-shorter/internal/model"
	"gorm.io/gorm"
)

type URLRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) Save(url *model.URL) (*model.URL, error) {
	err := r.db.Create(url).Error
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (r *URLRepository) FindByShortenCode(code string) (*model.URL, error) {
	var url model.URL
	result := r.db.Where("short_code = ?", code).First(&url)
	return &url, result.Error
}
