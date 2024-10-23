package repository

import (
	"context"

	"backend/app/domain/object"

	"gorm.io/gorm"
)

type CheckInEventDTO struct {
        StoreID       int
        StoreName     string
        StoreStaffID  int
}

type StudentRepository interface {
	Create(ctx context.Context, tx *gorm.DB, acc *object.Student) error
	FindByStudentId(ctx context.Context, studentId int) (*object.Student, error)
        CheckInEvent(ctx context.Context, event_id int, student_id int) (*CheckInEventDTO, error)
}
