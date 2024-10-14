package usecase

import (
	"backend/app/domain/object"
	"backend/app/domain/repository"
	"context"

	"gorm.io/gorm"
)

type OrderItem interface {
	Create(ctx context.Context, orderID, itemID, quantity int, arranges *string) (*CreateOrderItemDTO, error)
	GetByOrderId(ctx context.Context, orderID string) ([]*GetOrderItemDTO, error)
}

type CreateOrderItemDTO struct {
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
