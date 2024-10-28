package usecase

import (
	"backend/app/domain/repository"
	"context"

	"gorm.io/gorm"
)

type Report interface {
	SendEmail(ctx context.Context) (*CreateReportDTO, error)
}

type CreateReportDTO struct {
	IsSucceeded *bool
}

type report struct {
	db         *gorm.DB
	reportRepo repository.ReportRepository
}

var _ Report = (*report)(nil)

func NewReport(db *gorm.DB, reportRepo repository.ReportRepository) *report {
	return &report{
		db:         db,
		reportRepo: reportRepo,
	}
}

func (a *report) SendEmail(ctx context.Context) (*CreateReportDTO, error) {

	tx := a.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	if err := a.reportRepo.SendEmail(ctx, tx); err != nil {
		return nil, err
	}

	return &CreateReportDTO{}, nil
}
