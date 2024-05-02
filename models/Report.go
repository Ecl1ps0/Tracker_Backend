package models

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	Logs              []byte `binding:"required"`
	Report            []byte `binding:"required"`
	Clipboard         []byte `binding:"required"`
	StudentSolutionID uint
	StudentSolution   StudentSolution `gorm:"foreignKey:StudentSolutionID"`
}
