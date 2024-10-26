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

func (r *OrderItemRepositoryImpl) GetByID(ctx context.Context, id string) (*object.OrderItem, error) {
	var orderItem object.OrderItem

	if err := r.db.WithContext(ctx).Where("id = ?", id).Preload("Item").First(&orderItem).Error; err != nil {
		return nil, fmt.Errorf("failed to find order item by id: %w", err)
	}

	return &orderItem, nil
}

func (r *OrderItemRepositoryImpl) GetByOrderID(ctx context.Context, orderID string) ([]*object.OrderItem, error) {
	var orderItems []*object.OrderItem
	if err := r.db.WithContext(ctx).Where("order_id = ?", orderID).Preload("Item").Find(&orderItems).Error; err != nil {
		return nil, fmt.Errorf("failed to find order items by order id: %w", err)
	}
	return orderItems, nil
}

func (r *OrderItemRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, orderItem *object.OrderItem) error {
	if err := tx.WithContext(ctx).Model(&object.OrderItem{}).Where("id = ?", orderItem.ID).Updates(orderItem).Error; err != nil {
		return fmt.Errorf("failed to update order item: %w", err)
	}
	return nil
}

func (r *OrderItemRepositoryImpl) Delete(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&object.OrderItem{}).Error; err != nil {
		return fmt.Errorf("failed to delete order item: %w", err)
	}
	return nil
}
