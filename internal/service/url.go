package service

import "github.com/lukestarwalk/url-shorter/internal/repository"

type ULRService struct {
	repository *repository.URLRepository
}

func NewService(repository *repository.URLRepository) *ULRService {
	return &ULRService{repository: repository}
}
