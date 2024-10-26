package repository

import (
	"context"

	"backend/app/domain/object"

	"gorm.io/gorm"
)

type OrderItemRepository interface {
	Create(ctx context.Context, tx *gorm.DB, orderItem *object.OrderItem) error
	GetByID(ctx context.Context, id string) (*object.OrderItem, error)
	GetByOrderID(ctx context.Context, orderID string) ([]*object.OrderItem, error)
	Update(ctx context.Context, tx *gorm.DB, order *object.OrderItem) error
	Delete(ctx context.Context, id string) error
}
