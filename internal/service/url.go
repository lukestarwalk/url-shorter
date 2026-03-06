package service

import (
	"crypto/rand"
	"time"

	"github.com/google/uuid"
	"github.com/lukestarwalk/url-shorter/internal/model"
	"github.com/lukestarwalk/url-shorter/internal/repository"
)

type ULRService struct {
	repository *repository.URLRepository
}

func NewService(repository *repository.URLRepository) *ULRService {
	return &ULRService{repository: repository}
}

func (us *ULRService) Save(originalURL string) (*model.URL,error) {
	url := &model.URL{
		ID:          uuid.New(),
		ShortCode:   generateCode(),
		OriginalURL: originalURL,
		CreatedAt:   time.Now(),
	}
	return us.repository.Save(url)
}

func (us *ULRService) FindByShortenCode(code string) (*model.URL, error) {
	return us.repository.FindByShortenCode(code)
}

func generateCode() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 6)
	rand.Read(b)
	for i := range b {
		b[i] = chars[b[i]%byte(len(chars))]
	}
	return string(b)
}
