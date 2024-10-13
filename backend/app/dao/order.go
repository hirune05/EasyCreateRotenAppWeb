package dao

import (
	"backend/app/domain/object"
	"backend/app/domain/repository"
	"context"
	"fmt"

	"gorm.io/gorm"
)

var _ repository.OrderRepository = (*OrderRepositoryImpl)(nil)

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{db: db}
}

func (r *OrderRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, order *object.Order) error {
	if err := tx.WithContext(ctx).Create(order).Error; err != nil {
		return fmt.Errorf("failed to insert order: %w", err)
	}
	return nil
}
