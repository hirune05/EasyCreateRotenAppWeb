package usecase

import (
	"backend/app/domain/object"
	"backend/app/domain/repository"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type OrderItem interface {
	Create(ctx context.Context, orderID, itemID, quantity int, arranges *string) (*CreateOrderItemDTO, error)
	Update(ctx context.Context, id string, quantity int, arranges *string) (*UpdateOrderItemDTO, error)
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*GetOrderItemDTO, error)
	GetByOrderId(ctx context.Context, orderID string) ([]*GetOrderItemDTO, error)
}

type CreateOrderItemDTO struct {
	OrderItem *object.OrderItem
}
type UpdateOrderItemDTO struct {
	OrderItem *object.OrderItem
}
type GetOrderItemDTO struct {
	OrderItem *object.OrderItem
}

type orderItem struct {
	db            *gorm.DB
	orderItemRepo repository.OrderItemRepository
}

var _ OrderItem = (*orderItem)(nil)

func NewOrderItem(db *gorm.DB, orderItemRepo repository.OrderItemRepository) *orderItem {
	return &orderItem{
		db:            db,
		orderItemRepo: orderItemRepo,
	}
}

func (a *orderItem) Create(ctx context.Context, orderID, itemID, quantity int, arranges *string) (*CreateOrderItemDTO, error) {
	orderItem, err := object.NewOrderItem(orderID, itemID, quantity, arranges)
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

	if err := a.orderItemRepo.Create(ctx, tx, orderItem); err != nil {
		return nil, err
	}

	return &CreateOrderItemDTO{
		OrderItem: orderItem,
	}, nil
}

func (u *orderItem) GetByOrderId(ctx context.Context, orderID string) ([]*GetOrderItemDTO, error) {
	orderItems, err := u.orderItemRepo.GetByOrderId(ctx, orderID)
	if err != nil {
		return nil, err
	}

	var orderItemDTOs []*GetOrderItemDTO
	for _, item := range orderItems {
		orderItemDTOs = append(orderItemDTOs, &GetOrderItemDTO{
			OrderItem: &object.OrderItem{
				ID:        item.ID,
				OrderID:   item.OrderID,
				ItemID:    item.ItemID,
				Quantity:  item.Quantity,
				Arranges:  item.Arranges,
				CreatedAt: item.CreatedAt,
				UpdatedAt: item.UpdatedAt,
			},
		})
	}

	return orderItemDTOs, nil
}

func (a *orderItem) GetByID(ctx context.Context, id string) (*GetOrderItemDTO, error) {
	orderItem, err := a.orderItemRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get order by id: %w", err)
	}

	return &GetOrderItemDTO{
		OrderItem: orderItem,
	}, nil
}

func (a *orderItem) Update(ctx context.Context, id string, quantity int, arranges *string) (*UpdateOrderItemDTO, error) {
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

	orderItem, err := a.orderItemRepo.GetByID(ctx, id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	orderItem.Quantity = quantity
	orderItem.Arranges = arranges

	if err := a.orderItemRepo.Update(ctx, tx, orderItem); err != nil {
		tx.Rollback()
		return nil, err
	}

	return &UpdateOrderItemDTO{
		OrderItem: orderItem,
	}, nil
}

func (a *orderItem) Delete(ctx context.Context, id string) error {
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

	if err := a.orderItemRepo.Delete(ctx, id); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete order: %w", err)
	}

	return nil
}
