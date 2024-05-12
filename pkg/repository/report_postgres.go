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

func (r *ReportPostgres) GetAllReports() ([]models.Report, error) {
	var reports []models.Report
	if result := r.db.Find(&reports); result.Error != nil {
		return nil, result.Error
	}

	return reports, nil
}

func (r *ReportPostgres) CreateReport(data models.Report) (uint, error) {
	if result := r.db.Create(&data); result.Error != nil {
		return 0, nil
	}

	return data.ID, nil
}
