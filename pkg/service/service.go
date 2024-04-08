package service

import (
	"Proctor/models"
	"Proctor/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (uint, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (uint, error)
}

type FileHandler interface {
	SaveFile(data models.Data) (uint, error)
}

type Service struct {
	Authorization
	FileHandler
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		FileHandler:   NewFileHandlerService(repos),
	}
}
