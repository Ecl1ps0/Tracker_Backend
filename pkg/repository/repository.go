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
	CreateReport(data models.Report) (uint, error)
}

type Role interface {
	GetAll() ([]models.UserRole, int64, error)
	CreateDefault() ([]models.UserRole, error)
}

type User interface {
	GetProfile(id uint) (models.User, error)
	GetRoleByID(id uint) (uint, error)
}

type Task interface {
	CreateTask(task models.Task) (uint, error)
	GetTaskByID(id uint) (models.Task, error)
	DeleteTask(id uint) error
}

type Repository struct {
	Authorization
	FileHandler
	Role
	User
	Task
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		FileHandler:   NewFileHandlerPostgres(db),
		Role:          NewRolePostgres(db),
		User:          NewUserPostgres(db),
		Task:          NewTaskPostgres(db),
	}
}
