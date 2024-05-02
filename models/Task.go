package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Title       string `binding:"required"`
	Description string `binding:"required"`
	AccessFrom  *time.Time
	AccessTo    *time.Time
	TeacherID   uint
	Teacher     *User `gorm:"foreignKey:TeacherID"`
}
