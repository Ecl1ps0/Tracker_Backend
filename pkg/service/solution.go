package service

import (
	"Proctor/models"
	"Proctor/pkg/repository"
	"github.com/shopspring/decimal"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"regexp"
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

func (r *SolutionService) GetSolutionByID(id uint) (models.StudentSolution, error) {
	return r.repo.GetSolutionByID(id)
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

func (r *SolutionService) UpdateCheatingRate(id uint, rate decimal.Decimal) error {
	return r.repo.UpdateCheatingRate(id, rate)
}

func (r *SolutionService) UpdateFinalGrade(id uint, grade decimal.Decimal) error {
	return r.repo.UpdateFinalGrade(id, grade)
}

func (r *SolutionService) GetTeacherBySolutionID(id uint) (uint, error) {
	return r.repo.GetTeacherBySolutionID(id)
}

func (r *SolutionService) GenerateCheatingRate(solution string) (decimal.Decimal, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return decimal.NewFromFloat(-1), err
	}

	cmd := exec.Command("C:/Program Files/Git/bin/bash.exe", pwd+"/run_model.sh", viper.GetString("info"))
	cmd.Env = append(os.Environ(), "SOLUTION="+solution)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return decimal.NewFromFloat(-1), err
	}

	re := regexp.MustCompile(`"Prediction": ([0-9]+(?:\.[0-9]+)?)`)
	rate, err := decimal.NewFromString(string(re.FindSubmatch(output)[1]))
	if err != nil {
		return decimal.NewFromFloat(-1), err
	}

	return rate, nil
}
