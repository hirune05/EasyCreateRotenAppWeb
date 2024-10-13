package usecase

import (
	"backend/app/domain/object"
	"backend/app/domain/repository"
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Order interface {
	Create(ctx context.Context, storeID, storeStaffID, status int, pickedUpAt *time.Time) (*CreateOrderDTO, error)
	Update(ctx context.Context, id string, storeID, storeStaffID, status int, pickedUpAt *time.Time) (*UpdateOrderDTO, error)
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*GetOrderDTO, error)
	GetAll(ctx context.Context) ([]*GetOrderDTO, error)
}
type order struct {
	db        *gorm.DB
	orderRepo repository.OrderRepository
}

type CreateOrderDTO struct {
	Order *object.Order
}

type UpdateOrderDTO struct {
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

func (a *order) Update(ctx context.Context, id string, storeID, storeStaffID, status int, pickedUpAt *time.Time) (*UpdateOrderDTO, error) {
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

	order, err := a.orderRepo.GetByID(ctx, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	order.StoreID = storeID
	order.StoreStaffID = storeStaffID
	order.Status = status
	order.PickedUpAt = pickedUpAt
	order.UpdatedAt = time.Now()

	if err := a.orderRepo.Update(ctx, tx, order); err != nil {
		tx.Rollback()
		return nil, err
	}

	return &UpdateOrderDTO{
		Order: order,
	}, nil
}

func (a *order) Delete(ctx context.Context, id string) error {
	tx := a.db.Begin()
	if tx.Error != nil {
		return tx.Error
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

	if err := a.orderRepo.Delete(ctx, id); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete order: %w", err)
	}

	return nil
}

func (a *order) GetByID(ctx context.Context, id string) (*GetOrderDTO, error) {
	order, err := a.orderRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get order by id: %w", err)
	}

	return &GetOrderDTO{
		Order: order,
	}, nil
}

func (r *order) GetAll(ctx context.Context) ([]*GetOrderDTO, error) {
	var orders []*object.Order

	if err := r.db.WithContext(ctx).Find(&orders).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve all orders: %w", err)
	}

	var orderDTOs []*GetOrderDTO
	for _, order := range orders {
		orderDTO := &GetOrderDTO{
			Order: order,
		}
		orderDTOs = append(orderDTOs, orderDTO)
	}

	return orderDTOs, nil
}
