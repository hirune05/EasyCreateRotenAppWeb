package usecase

import (
	"backend/app/domain/object"
	"backend/app/domain/repository"
	"context"

	"gorm.io/gorm"
)

type AdminUser interface {
	Create(ctx context.Context, username, email, password string) (*CreateAdminUserDTO, error)
}

type CreateAdminUserDTO struct {
	AdminUser *object.AdminUser
}

type adminUser struct {
	db            *gorm.DB
	adminUserRepo repository.AdminUserRepository
}

var _ AdminUser = (*adminUser)(nil)

func NewAdminUser(db *gorm.DB, adminUserRepo repository.AdminUserRepository) *adminUser {
	return &adminUser{
		db:            db,
		adminUserRepo: adminUserRepo,
	}
}

func (a *adminUser) Create(ctx context.Context, username, email, password string) (*CreateAdminUserDTO, error) {
	adminUser, err := object.Create(username, email, password)
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

	if err := a.adminUserRepo.Create(ctx, tx, adminUser); err != nil {
		return nil, err
	}

	return &CreateAdminUserDTO{
		AdminUser: adminUser,
	}, nil
}
