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

func (r *OrderRepositoryImpl) GetByID(ctx context.Context, id string) (*object.Order, error) {
	var order object.Order

	if err := r.db.WithContext(ctx).
                        Where("id = ?", id).
                        Preload("OrderItems.Item").
                        Preload("StoreStaff.Student", func(db *gorm.DB) *gorm.DB { return db.Select("id, name") }).
                        First(&order).
                        Error; err != nil {
		return nil, fmt.Errorf("failed to find order by id: %w", err)
	}

	return &order, nil
}

func (r *OrderRepositoryImpl) GetByStoreID(ctx context.Context, id string) ([]*object.Order, error) {
	var orders []*object.Order
	if err := r.db.WithContext(ctx).
                        Where("store_id = ?", id).
                        Preload("OrderItems.Item").
                        Preload("StoreStaff.Student", func(db *gorm.DB) *gorm.DB { return db.Select("id, name") }).
                        Find(&orders).
                        Error; err != nil {
		return nil, fmt.Errorf("failed to find order by store id: %w", err)
	}
	return orders, nil
}

func (r *OrderRepositoryImpl) GetByStatus(ctx context.Context, storeID, status string) ([]*object.Order, error) {
	var orders []*object.Order
	if err := r.db.WithContext(ctx).
                        Where("store_id = ? AND status = ?", storeID, status).
                        Preload("OrderItems.Item").
                        Preload("StoreStaff.Student", func(db *gorm.DB) *gorm.DB { return db.Select("id, name") }).
                        Find(&orders).Error; err != nil {
		return nil, fmt.Errorf("failed to find order by store id and status: %w", err)
	}
	return orders, nil
}

func (r *OrderRepositoryImpl) GetAll(ctx context.Context) ([]*object.Order, error) {
	var orders []*object.Order

	if err := r.db.WithContext(ctx).
                        Preload("OrderItems.Item").
                        Preload("StoreStaff.Student", func(db *gorm.DB) *gorm.DB { return db.Select("id, name") }).
                        Find(&orders).
                        Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve all orders: %w", err)
	}

	return orders, nil
}

func (r *OrderRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, order *object.Order) error {
	if err := tx.WithContext(ctx).Model(&object.Order{}).Where("id = ?", order.ID).Updates(order).Error; err != nil {
		return fmt.Errorf("failed to update order: %w", err)
	}
	return nil
}

func (r *OrderRepositoryImpl) Delete(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&object.Order{}).Error; err != nil {
		return fmt.Errorf("failed to delete order: %w", err)
	}
	return nil
}
