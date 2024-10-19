package dao

import (
	"backend/app/domain/object"
	"backend/app/domain/repository"
	"context"
	"fmt"

	"gorm.io/gorm"
)

var _ repository.StoreStaffRepository = (*StoreStaffRepositoryImpl)(nil)

type StoreStaffRepositoryImpl struct {
	db *gorm.DB
}

func NewStoreStaffRepository(db *gorm.DB) *StoreStaffRepositoryImpl {
	return &StoreStaffRepositoryImpl{db: db}
}

func (r *StoreStaffRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, storeStaff *object.StoreStaff) error {
	if err := tx.WithContext(ctx).Create(storeStaff).Error; err != nil {
		return fmt.Errorf("failed to insert admin user: %w", err)
	}
	return nil
}

func (r *StoreStaffRepositoryImpl) FindByStudentNumber(ctx context.Context, student_number int) (*object.StoreStaff, error) {
	var storeStaff object.StoreStaff

	if err := r.db.WithContext(ctx).Where("student_number = ?", student_number).First(&storeStaff).Error; err != nil {
		return nil, fmt.Errorf("failed to find store_staff by student_number: %w", err)
	}

	return &storeStaff, nil
}
