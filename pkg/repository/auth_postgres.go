package repository

import (
	"Proctor/models"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (uint, error) {
	if result := r.db.Create(&user); result.Error != nil {
		return 0, result.Error
	}

	return user.ID, nil
}

func (r *AuthPostgres) GetUser(email, password string) (models.User, error) {
	var user models.User

	result := r.db.Where("email = ? AND password = ?", email, password).First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}

	return user, nil
}
