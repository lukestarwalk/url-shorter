package repository

import (
	"github.com/lukestarwalk/url-shorter/internal/model"
	"gorm.io/gorm"
)

type URL model.URL

type URLRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) Save(url *URL) error {
	return r.db.Create(url).Error
}

func (r *URLRepository) FindByShortenCode(code string) (*URL, error) {
	var url URL
	result := r.db.Where("shortcode = ?", code).First(&url)
	return &url, result.Error
}
