package repository

import (
	"Proctor/models"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

const ID_SELECT_PARAM = "id=?"

type SolutionPostgres struct {
	db *gorm.DB
}

func NewSolutionPostgres(db *gorm.DB) *SolutionPostgres {
	return &SolutionPostgres{db: db}
}

func (r *SolutionPostgres) GetAllSolutions() ([]models.StudentSolution, error) {
	var solutions []models.StudentSolution
	if result := r.db.Preload("Report").Find(&solutions); result.Error != nil {
		return nil, result.Error
	}

	return solutions, nil
}

func (r *SolutionPostgres) GetSolutionsByStudentID(id uint) ([]models.StudentSolution, error) {
	var solutions []models.StudentSolution
	if result := r.db.Table("student_solutions").
		Joins("join student_tasks on student_solutions.student_task_id = student_tasks.id").
		Joins("join users on student_tasks.student_id = users.id").
		Where("users.id = ?", id).Find(&solutions); result.Error != nil {
		return nil, result.Error
	}

	return solutions, nil
}

func (r *SolutionPostgres) GetSolutionByID(id uint) (models.StudentSolution, error) {
	var solution models.StudentSolution
	if result := r.db.Preload("Report").Where(ID_SELECT_PARAM, id).Find(&solution); result.Error != nil {
		return models.StudentSolution{}, result.Error
	}

	return solution, nil
}

func (r *SolutionPostgres) GetUserSolutionsOnSolvedTask(id uint) ([]models.StudentSolution, error) {
	var users []models.StudentSolution
	if results := r.db.Table("student_solutions").
		Select("student_solutions.*").
		Joins("join student_tasks on student_solutions.student_task_id = student_tasks.id").
		Joins("join users on student_tasks.student_id = users.id").
		Where("student_tasks.task_id = ?", id).
		Scan(&users); results.Error != nil {
		return nil, results.Error
	}

	return users, nil
}

func (r *SolutionPostgres) GetStudentSolutionOnTask(studentSolutionId uint) (models.StudentSolution, error) {
	var solution models.StudentSolution
	if result := r.db.Where("student_task_id = ?", studentSolutionId).Find(&solution); result.Error != nil {
		return models.StudentSolution{}, result.Error
	}

	return solution, nil
}

func (r *SolutionPostgres) CreateSolution(solution models.StudentSolution) (uint, error) {
	if result := r.db.Create(&solution); result.Error != nil {
		return 0, result.Error
	}

	return solution.ID, nil
}

func (r *SolutionPostgres) UpdateCheatingRate(id uint, rate decimal.Decimal) error {
	if result := r.db.Model(&models.StudentSolution{}).Where(ID_SELECT_PARAM, id).Update("CheatingResult", rate); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *SolutionPostgres) UpdateFinalGrade(id uint, grade decimal.Decimal) error {
	if result := r.db.Model(&models.StudentSolution{}).Where(ID_SELECT_PARAM, id).Update("FinalGrade", grade); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *SolutionPostgres) GetTeacherBySolutionID(id uint) (uint, error) {
	var teacherId uint
	if result := r.db.Table("student_solutions").
		Select("tasks.teacher_id").
		Joins("join student_tasks on student_tasks.id = student_solutions.student_task_id").
		Joins("join tasks on tasks.id = student_tasks.task_id").
		Where("student_solutions.id = ?", id).Find(&teacherId); result.Error != nil {
		return 0, result.Error
	}

	return teacherId, nil
}
