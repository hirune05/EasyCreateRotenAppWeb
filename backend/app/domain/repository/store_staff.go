package repository

import (
	"context"

	"backend/app/domain/object"

	"gorm.io/gorm"
)

type StoreStaffRepository interface {
	Create(ctx context.Context, tx *gorm.DB, acc *object.StoreStaff) error
}
