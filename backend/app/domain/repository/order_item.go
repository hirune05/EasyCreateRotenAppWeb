package repository

import (
	"context"

	"backend/app/domain/object"

	"gorm.io/gorm"
)

type OrderItemRepository interface {
	Create(ctx context.Context, tx *gorm.DB, orderItem *object.OrderItem) error
	GetByOrderId(ctx context.Context, orderID string) ([]*object.OrderItem, error)
}
