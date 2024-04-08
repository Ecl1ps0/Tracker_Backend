package repository

import (
	"Proctor/models"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user models.User) (uint, error)
	GetUser(email, password string) (models.User, error)
}

type FileHandler interface {
	Create(data models.Data) (uint, error)
}

type Repository struct {
	Authorization
	FileHandler
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		FileHandler:   NewFileHandlerPostgres(db),
	}
}
