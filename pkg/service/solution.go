package service

import (
	"Proctor/models"
	"Proctor/pkg/repository"
)

type SolutionService struct {
	repo repository.Solution
}

func NewSolutionService(repo repository.Solution) *SolutionService {
	return &SolutionService{repo: repo}
}

func (r *SolutionService) CreateSolution(solution models.StudentSolution) (uint, error) {
	return r.repo.CreateSolution(solution)
}
