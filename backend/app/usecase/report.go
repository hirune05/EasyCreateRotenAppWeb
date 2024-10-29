package usecase

import (
	"backend/app/domain/object"
	"backend/app/domain/repository"
	"context"

	"gorm.io/gorm"
)

type Report interface {
        Create(ctx context.Context, storeID *int, storeStaffID *int, title string, description string) (*ReportDTO, error)
	SendEmail(ctx context.Context, subject string, body string) (*CreateReportDTO, error)
}
type report struct {
	db        *gorm.DB
	reportRepo repository.ReportRepository
}

type ReportDTO struct {
	Report *object.Report
}

type CreateReportDTO struct {
	Message string
}

var _ Report = (*report)(nil)

func NewReport(db *gorm.DB, reportRepo repository.ReportRepository) *report {
	return &report{
		db:         db,
		reportRepo: reportRepo,
	}
}

func (a *report) Create(ctx context.Context, storeID *int, storeStaffID *int, title string, description string) (*ReportDTO, error) {
	acc, err := object.NewReport(storeID, storeStaffID, title, description)
	if err != nil {
		return nil, err
	}

	tx := a.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err := a.reportRepo.Create(ctx, tx, acc); err != nil {
		return nil, err
	}

        _, err2 := a.reportRepo.SendEmail(ctx, tx, title, description)
        if err2 != nil {
		return nil, err2
        }

	return &ReportDTO{
		Report: acc,
	}, nil
}
func (a *report) SendEmail(ctx context.Context, subject string, body string) (*CreateReportDTO, error) {

	tx := a.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	message, err := a.reportRepo.SendEmail(ctx, tx, subject, body)
	if err != nil {
		return nil, err
	}

	return &CreateReportDTO{Message: message}, nil
}
