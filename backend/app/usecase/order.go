package usecase

import (
	"backend/app/domain/object"
	"backend/app/domain/repository"
	"context"
	"time"

	"gorm.io/gorm"
)

type Order interface {
	Create(ctx context.Context, storeID, storeStaffID, status int, pickedUpAt *time.Time) (*CreateOrderDTO, error)
}

type order struct {
	db        *gorm.DB
	orderRepo repository.OrderRepository
}

type CreateOrderDTO struct {
	Order *object.Order
}

type GetOrderDTO struct {
	Order *object.Order
}

var _ Order = (*order)(nil)

func NewOrder(db *gorm.DB, orderRepo repository.OrderRepository) *order {
	return &order{
		db:        db,
		orderRepo: orderRepo,
	}
}

func (a *order) Create(ctx context.Context, storeID, storeStaffID, status int, pickedUpAt *time.Time) (*CreateOrderDTO, error) {
	acc, err := object.NewOrder(storeID, storeStaffID, status, pickedUpAt)
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

	if err := a.orderRepo.Create(ctx, tx, acc); err != nil {
		return nil, err
	}

	return &CreateOrderDTO{
		Order: acc,
	}, nil
}
