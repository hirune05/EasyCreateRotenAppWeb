package repository

import (
	"context"

	"backend/app/domain/object"
	"gorm.io/gorm"
)

type ReportRepository interface {
	Create(ctx context.Context, tx *gorm.DB, report *object.Report) error
	SendEmail(ctx context.Context, tx *gorm.DB, subject string, body string) (string, error)
}
