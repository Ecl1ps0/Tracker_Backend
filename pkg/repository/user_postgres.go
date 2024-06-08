package repository

import (
	"Proctor/models"
	"errors"
	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if result := r.db.Find(&users); result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (r *UserPostgres) GetProfile(id uint) (models.User, error) {
	var user models.User
	result := r.db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}

func (r *UserPostgres) GetAllStudents() ([]models.User, error) {
	var students []models.User
	if result := r.db.Where("role_id = ?", 1).Find(&students); result.Error != nil {
		return nil, result.Error
	}

	return students, nil
}

func (r *UserPostgres) GetStudentBySolutionID(id uint) (models.User, error) {
	var student models.User
	if result := r.db.Table("student_solutions").
		Select("users.*").
		Joins("join student_tasks on student_solutions.student_task_id = student_tasks.id").
		Joins("join users on student_tasks.student_id = users.id").
		Where("student_solutions.id = ?", id).Scan(&student); result.Error != nil {
		return models.User{}, result.Error
	}

	return student, nil
}

func (r *UserPostgres) GetStudentsByTeacherID(id uint) ([]models.User, error) {
	var students []models.User
	if result := r.db.Table("student_tasks").
		Select("users.*").
		Joins("join tasks on tasks.id = student_tasks.task_id").
		Joins("join users on users.id = student_tasks.student_id").
		Where("tasks.teacher_id = ?", id).Scan(&students); result.Error != nil {
		return nil, result.Error
	}

	return students, nil
}

func (r *UserPostgres) AddStudentToTask(studentTask models.StudentTask) error {
	var existStudentTask models.StudentTask
	if err := r.db.Table("student_task").Where("task_id = ? and student_id = ?", studentTask.TaskID, studentTask.StudentID).First(&existStudentTask).Error; err != nil {
		return errors.New("duplicate entry")
	}

	if result := r.db.Create(&studentTask); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *UserPostgres) GetRoleByID(id uint) (uint, error) {
	var user models.User
	result := r.db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return 0, result.Error
	}

	return user.RoleID, nil
}
