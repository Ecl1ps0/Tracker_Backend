package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type StudentSolution struct {
	gorm.Model
	Solution       string `binding:"required"`
	TimeStart      *time.Time
	TimeEnd        *time.Time
	CheatingResult decimal.Decimal `sql:"type:decimal(5,2);default:0;"`
	StudentTaskID  uint
	StudentTask    *StudentTask `gorm:"foreignKey:StudentTaskID"`
	ReportID       uint
	Report         *Report `gorm:"foreignKey:ReportID"`
}
