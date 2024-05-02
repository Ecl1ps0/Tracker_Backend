package models

import "gorm.io/gorm"

type StudentTask struct {
	gorm.Model
	StudentID uint
	Student   *User `gorm:"foreignKey:StudentID"`
	TaskID    uint
	Task      *Task `gorm:"foreignKey:TaskID"`
}
