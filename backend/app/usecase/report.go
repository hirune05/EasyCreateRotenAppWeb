package usecase

import (
	"backend/app/domain/object"
	"backend/app/domain/repository"
	"context"

	"gorm.io/gorm"
)

type Report interface {
        Create(ctx context.Context, storeID *int, storeStaffID *int, description string) (*ReportDTO, error)
}
type report struct {
	db        *gorm.DB
	reportRepo repository.ReportRepository
}

type ReportDTO struct {
	Report *object.Report
}

var _ Report = (*report)(nil)

func NewReport(db *gorm.DB, reportRepo repository.ReportRepository) *report {
	return &report{
		db:        db,
		reportRepo: reportRepo,
	}
}

func (a *report) Create(ctx context.Context, storeID *int, storeStaffID *int, description string) (*ReportDTO, error) {
	acc, err := object.NewReport(storeID, storeStaffID, description)
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

	return &ReportDTO{
		Report: acc,
	}, nil
}
