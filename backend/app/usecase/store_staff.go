package usecase

import (
	"backend/app/domain/object"
	"backend/app/domain/repository"
	"context"

	"gorm.io/gorm"
)

type StoreStaff interface {
	Create(ctx context.Context, name string, password string, studentNumber int, role int, storeId *int) (*CreateStoreStaffDTO, error)
}

type CreateStoreStaffDTO struct {
	StoreStaff *object.StoreStaff
}

type storeStaff struct {
	db             *gorm.DB
	storeStaffRepo repository.StoreStaffRepository
}

var _ StoreStaff = (*storeStaff)(nil)

func NewStoreStaff(db *gorm.DB, storeStaffRepo repository.StoreStaffRepository) *storeStaff {
	return &storeStaff{
		db:             db,
		storeStaffRepo: storeStaffRepo,
	}
}

func (a *storeStaff) Create(ctx context.Context, name string, password string, studentNumber int, role int, storeId *int) (*CreateStoreStaffDTO, error) {
	storeStaff, err := object.NewStoreStaff(name, password, studentNumber, role, storeId)
	if err != nil {
		return nil, err
	}

	tx := a.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err := a.storeStaffRepo.Create(ctx, tx, storeStaff); err != nil {
		return nil, err
	}

	return &CreateStoreStaffDTO{
		StoreStaff: storeStaff,
	}, nil
}
