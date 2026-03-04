package service

import "github.com/lukestarwalk/url-shorter/internal/repository"

type URLRepository repository.URLRepository

type ULRService struct {
	repository *URLRepository
}

func NewService(repository *URLRepository) *ULRService{
	return &ULRService{repository: repository}
}
