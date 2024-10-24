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

func (r *StudentRepositoryImpl) CheckInEvent(ctx context.Context, event_id int, student_id int) (*repository.CheckInEventDTO, error) {
        var dto repository.CheckInEventDTO
        err := r.db.WithContext(ctx).Table("students").
                Select("stores.id AS StoreID, stores.name AS StoreName, store_staffs.id AS StoreStaffID").
                Joins("INNER JOIN store_staffs ON students.id = store_staffs.student_id").
                Joins("INNER JOIN stores ON store_staffs.store_id = stores.id").
                Joins("INNER JOIN events ON stores.event_id = events.id").
                Where("students.id = ? AND events.id = ?", student_id, event_id).
                Limit(1).Scan(&dto).Error


        if err != nil {
                return nil, err
        }
        return &dto, nil
}
