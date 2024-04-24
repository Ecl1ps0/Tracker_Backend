package repository

import (
	"Proctor/models"
	"gorm.io/gorm"
)

type RolePostgres struct {
	db *gorm.DB
}

func NewRolePostgres(db *gorm.DB) *RolePostgres {
	return &RolePostgres{db: db}
}

func (r *RolePostgres) GetAll() ([]models.UserRole, int64, error) {
	var roles []models.UserRole
	var count int64

	result := r.db.Find(&roles).Count(&count)
	if result.Error != nil {
		return []models.UserRole{}, 0, result.Error
	}

	return roles, count, nil
}

func (r *RolePostgres) CreateDefault() ([]models.UserRole, error) {
	roles := []models.UserRole{{Name: "Student"}, {Name: "Teacher"}, {Name: "Admin"}}

	result := r.db.Create(&roles)
	if result.Error != nil {
		return []models.UserRole{}, result.Error
	}

	return roles, nil
}
