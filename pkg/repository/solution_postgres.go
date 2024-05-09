package repository

import (
	"Proctor/models"
	"gorm.io/gorm"
)

type SolutionPostgres struct {
	db *gorm.DB
}

func NewSolutionPostgres(db *gorm.DB) *SolutionPostgres {
	return &SolutionPostgres{db: db}
}

func (r *SolutionPostgres) CreateSolution(solution models.StudentSolution) (uint, error) {
	if result := r.db.Create(&solution); result.Error != nil {
		return 0, result.Error
	}

	return solution.ID, nil
}
