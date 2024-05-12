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
	CheatingResult decimal.Decimal `sql:"type:decimal(5,2);check:(CheatingResult >= 0 AND CheatingResult <= 100);default:0;"`
	FinalGrade     decimal.Decimal `sql:"type:decimal(5,2);check:(FinalResult >= 0 AND FinalResult <= 100);default:0;"`
	StudentTaskID  uint
	StudentTask    *StudentTask `gorm:"foreignKey:StudentTaskID"`
	ReportID       uint
	Report         *Report `gorm:"foreignKey:ReportID"`
}
