package dao

import (
	"backend/app/domain/object"
	"backend/app/domain/repository"
	"context"
	"fmt"

	"gorm.io/gorm"
)

var _ repository.ReportRepository = (*ReportRepositoryImpl)(nil)

type ReportRepositoryImpl struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) *ReportRepositoryImpl {
	return &ReportRepositoryImpl{db: db}
}

func (r *ReportRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, report *object.Report) error {
	if err := tx.WithContext(ctx).Create(report).Error; err != nil {
		return fmt.Errorf("failed to insert report: %w", err)
	}
	return nil
}
