package service

import (
	"Proctor/models"
	"Proctor/pkg/repository"
)

type ReportService struct {
	repo repository.Report
}

func NewReportService(repo repository.Report) *ReportService {
	return &ReportService{repo: repo}
}

func (s *ReportService) CreateReport(data models.Report) (uint, error) {
	return s.repo.CreateReport(data)
}
