package repository

import (
	"Proctor/models"
	"gorm.io/gorm"
)

type ReportPostgres struct {
	db *gorm.DB
}

func NewReportPostgres(db *gorm.DB) *ReportPostgres {
	return &ReportPostgres{db: db}
}

func (r *ReportPostgres) CreateReport(data models.Report) (uint, error) {
	if result := r.db.Create(&data); result.Error != nil {
		return 0, nil
	}

	return data.ID, nil
}
