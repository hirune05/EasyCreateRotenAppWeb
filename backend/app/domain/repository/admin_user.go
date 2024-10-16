package repository

import (
	"context"

	"backend/app/domain/object"

	"gorm.io/gorm"
)

type AdminUserRepository interface {
	Create(ctx context.Context, tx *gorm.DB, acc *object.AdminUser) error
}
