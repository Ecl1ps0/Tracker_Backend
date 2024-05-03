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
	SaveFile(data models.Report) (uint, error)
}

type Role interface {
	SetDefaultRoles() ([]models.UserRole, error)
}

type User interface {
	GetProfile(userId uint) (models.User, error)
	AddStudentToTask(studentId, taskId uint) error
	GetRoleByUserID(id uint) (uint, error)
}

type Task interface {
	CreateTask(task models.Task) (uint, error)
	GetAllTasks() ([]models.Task, error)
	GetTaskByID(id uint) (models.Task, error)
	GetAllTeacherTasks(id uint) ([]models.Task, error)
	GetAllStudentTasks(id uint) ([]models.Task, error)
	UpdateTask(task models.Task) error
	DeleteTask(id uint) error
}

type Service struct {
	Authorization
	FileHandler
	Role
	User
	Task
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		FileHandler:   NewFileHandlerService(repos),
		Role:          NewRoleService(repos),
		User:          NewUserService(repos),
		Task:          NewTaskService(repos),
	}
}
