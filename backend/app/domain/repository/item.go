package repository

import (
	"context"

	"backend/app/domain/object"

	"gorm.io/gorm"
)

type GetByStoreIDDTO struct {
        StoreID       int             `json:"store_id"`
        StoreName     string          `json:"store_name"`
        Items         []*object.Item  `json:"items"`
}

type ItemRepository interface {
	Create(ctx context.Context, tx *gorm.DB, item *object.Item) error
	GetByID(ctx context.Context, id string) (*object.Item, error)
	GetByStoreID(ctx context.Context, id string) (*GetByStoreIDDTO, error)
}
