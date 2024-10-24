package dao

import (
	"backend/app/domain/object"
	"backend/app/domain/repository"
	"context"
	"fmt"

	"gorm.io/gorm"
)

var _ repository.ItemRepository = (*ItemRepositoryImpl)(nil)

type ItemRepositoryImpl struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepositoryImpl {
	return &ItemRepositoryImpl{db: db}
}

func (r *ItemRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, item *object.Item) error {
	if err := tx.WithContext(ctx).Create(item).Error; err != nil {
		return fmt.Errorf("failed to insert item: %w", err)
	}
	return nil
}

func (r *ItemRepositoryImpl) GetByID(ctx context.Context, id string) (*object.Item, error) {
	var item object.Item

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&item).Error; err != nil {
		return nil, fmt.Errorf("failed to find item by id: %w", err)
	}

	return &item, nil
}

func (r *ItemRepositoryImpl) GetByStoreID(ctx context.Context, id string) (*repository.GetByStoreIDDTO, error) {
	var items []*object.Item
	if err := r.db.WithContext(ctx).Table("items").Where("store_id = ?", id).Scan(&items).Error; err != nil {
		return nil, fmt.Errorf("failed to find item by store id: %w", err)
	}

        var store object.Store
	if err := r.db.WithContext(ctx).Table("stores").Where("id = ?", id).Scan(&store).Error; err != nil {
		return nil, fmt.Errorf("failed to find item by store id: %w", err)
	}

	return &repository.GetByStoreIDDTO{
          StoreID:    store.ID,
          StoreName:  store.Name,
          Items:  items,
        }, nil
}

