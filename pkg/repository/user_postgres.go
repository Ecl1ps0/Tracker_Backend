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

func (r *UserPostgres) GetProfile(id uint) (models.User, error) {
	var user models.User

	result := r.db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}
