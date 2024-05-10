package repository

import (
	"Proctor/models"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user models.User) (uint, error)
	GetUser(email, password string) (models.User, error)
}

type Report interface {
	CreateReport(data models.Report) (uint, error)
}

type Role interface {
	GetAll() ([]models.UserRole, int64, error)
	CreateDefault() ([]models.UserRole, error)
}

type User interface {
	GetProfile(id uint) (models.User, error)
	AddStudentToTask(studentTask models.StudentTask) error
	GetRoleByID(id uint) (uint, error)
}

type Task interface {
	CreateTask(task models.Task) (uint, error)
	GetAllTasks() ([]models.Task, error)
	GetTaskByID(id uint) (models.Task, error)
	GetAllTeacherTasks(id uint) ([]models.Task, error)
	GetAllStudentTasks(id uint) ([]models.StudentTask, error)
	UpdateTask(task models.Task) error
	DeleteTask(id uint) error
}

type Solution interface {
	CreateSolution(solution models.StudentSolution) (uint, error)
}

type Repository struct {
	Authorization
	Report
	Role
	User
	Task
	Solution
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Report:        NewReportPostgres(db),
		Role:          NewRolePostgres(db),
		User:          NewUserPostgres(db),
		Task:          NewTaskPostgres(db),
		Solution:      NewSolutionPostgres(db),
	}
}
