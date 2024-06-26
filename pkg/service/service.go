package service

import (
	"Proctor/models"
	"Proctor/models/DTO"
	"Proctor/pkg/repository"
	"github.com/shopspring/decimal"
)

type Authorization interface {
	CreateUser(user models.User) (uint, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (uint, error)
}

type Report interface {
	GetAllReports() ([]models.Report, error)
	CreateReport(data models.Report) (uint, error)
}

type Role interface {
	SetDefaultRoles() ([]models.UserRole, error)
}

type User interface {
	GetAllUsers() ([]DTO.UserDTO, error)
	GetProfile(userId uint) (DTO.UserDTO, error)
	GetAllStudents() ([]DTO.UserDTO, error)
	GetStudentBySolutionID(id uint) (DTO.UserDTO, error)
	GetStudentsByTeacherID(id uint) ([]DTO.UserDTO, error)
	AddStudentToTask(studentId, taskId uint) error
	GetRoleByUserID(id uint) (uint, error)
	UserToDTO(user models.User) DTO.UserDTO
	ParseUsersToDTOs(users []models.User) []DTO.UserDTO
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
	GetAllSolutions() ([]models.StudentSolution, error)
	GetSolutionsByStudentID(id uint) ([]models.StudentSolution, error)
	GetSolutionByID(id uint) (models.StudentSolution, error)
	GetUserSolutionsOnSolvedTask(id uint) ([]models.StudentSolution, error)
	GetStudentSolutionOnTask(studentSolutionId uint) (models.StudentSolution, error)
	CreateSolution(solution models.StudentSolution) (uint, error)
	UpdateCheatingRate(id uint, rate decimal.Decimal) error
	UpdateFinalGrade(id uint, grade decimal.Decimal) error
	GetTeacherBySolutionID(id uint) (uint, error)
	GenerateCheatingRate(solution string) (decimal.Decimal, error)
}

type Service struct {
	Authorization
	Report
	Role
	User
	Task
	Solution
	Redis *repository.RedisRepository
}

func NewService(repos *repository.Repository, redis *repository.RedisRepository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		Report:        NewReportService(repos),
		Role:          NewRoleService(repos),
		User:          NewUserService(repos),
		Task:          NewTaskService(repos),
		Solution:      NewSolutionService(repos),
		Redis:         redis,
	}
}
