package models

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	AssignmentID     uint
	Assignment       Task `gorm:"foreignKey:AssignmentID"`
	StudentID        uint
	Student          User `gorm:"foreignKey:StudentID"`
	AssignmentDataID uint
	AssignmentData   Data `gorm:"foreignKey:AssignmentDataID"`
}
