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

func (r *UserPostgres) GetProfile(id uint) (models.User, error) {
	var user models.User
	result := r.db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
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
