package service

import (
	"Proctor/models"
	"Proctor/models/DTO"
	"Proctor/pkg/repository"
	"github.com/shopspring/decimal"
)

type SolutionService struct {
	repo repository.Solution
}

func NewSolutionService(repo repository.Solution) *SolutionService {
	return &SolutionService{repo: repo}
}

func (r *SolutionService) GetAllSolutions() ([]models.StudentSolution, error) {
	return r.repo.GetAllSolutions()
}

func (r *SolutionService) GetUserSolutionsOnSolvedTask(id uint) ([]models.StudentSolution, error) {
	return r.repo.GetUserSolutionsOnSolvedTask(id)
}

func (r *SolutionService) GetStudentSolutionOnTask(studentSolutionId uint) (models.StudentSolution, error) {
	return r.repo.GetStudentSolutionOnTask(studentSolutionId)
}

func (r *SolutionService) CreateSolution(solution models.StudentSolution) (uint, error) {
	return r.repo.CreateSolution(solution)
}

func (r *SolutionService) UpdateCheatingRate(id uint, dto DTO.SolutionCheatingRateDTO) error {
	return r.repo.UpdateCheatingRate(id, dto.CheatingRate)
}

func (r *SolutionService) UpdateFinalGrade(id uint, grade decimal.Decimal) error {
	return r.repo.UpdateFinalGrade(id, grade)
}

func (r *SolutionService) GetTeacherBySolutionID(id uint) (uint, error) {
	return r.repo.GetTeacherBySolutionID(id)
}
