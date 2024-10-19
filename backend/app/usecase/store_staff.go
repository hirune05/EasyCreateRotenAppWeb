package usecase

import (
	"backend/app/domain/object"
	"backend/app/domain/repository"
	"context"
	"errors"
	"os"

	"gorm.io/gorm"
)

type StoreStaff interface {
	Create(ctx context.Context, name string, password string, studentNumber int, role int, storeId int) (*CreateStoreStaffDTO, error)
	Login(ctx context.Context, studentNumber int, password string) (*LoginStoreStaffDTO, error)
}
type LoginStoreStaffDTO struct {
	Token         string `json:"token"`
	UserID        int    `json:"user_id"`
	StudentNumber int    `json:"student_number"`
	Name          string `json:"name"`
	Role          int    `json:"role"`
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

func (a *storeStaff) Create(ctx context.Context, name string, password string, studentNumber int, role int, storeId int) (*CreateStoreStaffDTO, error) {
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

func (a *storeStaff) Login(ctx context.Context, studentNumber int, password string) (*LoginStoreStaffDTO, error) {
	storeStaff, err := a.storeStaffRepo.FindByStudentNumber(ctx, studentNumber)
	if err != nil {
		return nil, err
	}

	if !storeStaff.CheckPassword(password) {
		return nil, errors.New("invalid student number or password")
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return nil, errors.New("JWT secret key is not configured")
	}

	token, err := storeStaff.GenerateJWTToken(secretKey)
	if err != nil {
		return nil, err
	}

	return &LoginStoreStaffDTO{
		Token:         token,
		UserID:        storeStaff.ID,
		StudentNumber: storeStaff.StudentNumber,
		Name:          storeStaff.Name,
		Role:          storeStaff.Role,
	}, nil
}
