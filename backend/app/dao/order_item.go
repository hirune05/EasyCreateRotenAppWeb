package dao

import (
	"backend/app/domain/object"
	"backend/app/domain/repository"
	"context"
	"fmt"

	"gorm.io/gorm"
)

var _ repository.OrderItemRepository = (*OrderItemRepositoryImpl)(nil)

type OrderItemRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) *OrderItemRepositoryImpl {
	return &OrderItemRepositoryImpl{db: db}
}

func (r *OrderItemRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, orderItem *object.OrderItem) error {
	if err := tx.WithContext(ctx).Create(orderItem).Error; err != nil {
		return fmt.Errorf("failed to insert order item: %w", err)
	}
	return nil
}

func (r *OrderItemRepositoryImpl) GetByOrderId(ctx context.Context, orderID string) ([]*object.OrderItem, error) {
	var orderItems []*object.OrderItem
	if err := r.db.WithContext(ctx).Where("order_id = ?", orderID).Find(&orderItems).Error; err != nil {
		return nil, fmt.Errorf("failed to find order items by order_id: %w", err)
	}
	return orderItems, nil
}
