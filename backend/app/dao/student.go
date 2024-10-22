package dao

import (
	"backend/app/domain/object"
	"backend/app/domain/repository"
	"context"
	"fmt"

	"gorm.io/gorm"
)

var _ repository.StudentRepository = (*StudentRepositoryImpl)(nil)

type StudentRepositoryImpl struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepositoryImpl {
	return &StudentRepositoryImpl{db: db}
}

func (r *StudentRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, storeStaff *object.Student) error {
	if err := tx.WithContext(ctx).Create(storeStaff).Error; err != nil {
		return fmt.Errorf("failed to insert admin user: %w", err)
	}
	return nil
}

func (r *StudentRepositoryImpl) FindByStudentId(ctx context.Context, student_id int) (*object.Student, error) {
	var storeStaff object.Student

	if err := r.db.WithContext(ctx).Where("id = ?", student_id).First(&storeStaff).Error; err != nil {
		return nil, fmt.Errorf("failed to find store_staff by student_id: %w", err)
	}

	return &storeStaff, nil
}
