package repository

import (
	"Proctor/models"
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

func (r *UserPostgres) AddStudentToTask(studentTask models.StudentTask) error {
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
