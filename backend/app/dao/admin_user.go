package dao

import (
	"backend/app/domain/object"
	"backend/app/domain/repository"
	"context"
	"fmt"

	"gorm.io/gorm"
)

var _ repository.AdminUserRepository = (*AdminUserRepositoryImpl)(nil)

type AdminUserRepositoryImpl struct {
	db *gorm.DB
}

func NewAdminUserRepository(db *gorm.DB) *AdminUserRepositoryImpl {
	return &AdminUserRepositoryImpl{db: db}
}

func (r *AdminUserRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, adminUser *object.AdminUser) error {
	if err := tx.WithContext(ctx).Create(adminUser).Error; err != nil {
		return fmt.Errorf("failed to insert admin user: %w", err)
	}
	return nil
}
