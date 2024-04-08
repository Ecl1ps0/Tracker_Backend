package models

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	Name string `binding:"required"`
}
