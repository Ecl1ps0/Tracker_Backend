package service

import (
	"Proctor/models"
	"Proctor/pkg/repository"
)

type FileHandlerService struct {
	repo repository.FileHandler
}

func NewFileHandlerService(repo repository.FileHandler) *FileHandlerService {
	return &FileHandlerService{repo: repo}
}

func (s *FileHandlerService) SaveFile(data models.Report) (uint, error) {
	return s.repo.CreateReport(data)
}
