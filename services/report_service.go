package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
	"time"
)

type ReportService struct {
	repo repositories.ReportRepository
}

func NewReportService(repo repositories.ReportRepository) *ReportService {
	return &ReportService{
		repo: repo,
	}
}

func (s *ReportService) GetTodayReport() (*models.ReportResponse, error) {
	now := time.Now()

	start := time.Date(
		now.Year(), now.Month(), now.Day(),
		0, 0, 0, 0,
		now.Location(),
	)

	end := start.Add(24 * time.Hour)

	return s.repo.GetReport(start, end)
}

func (s *ReportService) GetReportByRange(start, end time.Time) (*models.ReportResponse, error) {
	return s.repo.GetReport(start, end)
}
