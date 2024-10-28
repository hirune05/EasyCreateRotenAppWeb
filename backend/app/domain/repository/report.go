package repository

import (
	"context"

	"gorm.io/gorm"
)

type ReportRepository interface {
	SendEmail(ctx context.Context, tx *gorm.DB) (string, error)
}
