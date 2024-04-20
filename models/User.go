package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Name     string `binding:"required"`
	Surname  *string
	Email    string `gorm:"unique" binding:"required"`
	RoleID   uint
	Role     *UserRole `gorm:"foreignKey:RoleID"`
	Password string    `binding:"required"`
	Tasks    []*Task   `gorm:"many2many:student_task;"`
}
