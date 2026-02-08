package repositories

import (
	"kasir-api/models"
	"time"
)

// ReportRepository defines contract for report operations
type ReportRepository interface {
	GetReport(start, end time.Time) (*models.ReportResponse, error)
}
