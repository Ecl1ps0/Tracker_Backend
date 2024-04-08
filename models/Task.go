package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `binding:"required"`
	Description string `binding:"required"`
	TeacherID   uint
	Teacher     User    `gorm:"foreignKey:TeacherID"`
	Students    []*User `gorm:"many2many:student_task;"`
}
