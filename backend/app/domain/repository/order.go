package repository

import (
	"context"

	"backend/app/domain/object"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(ctx context.Context, tx *gorm.DB, order *object.Order) error
	GetByID(ctx context.Context, id string) (*object.Order, error)
	GetAll(ctx context.Context) ([]*object.Order, error)
	Update(ctx context.Context, tx *gorm.DB, order *object.Order) error
	Delete(ctx context.Context, id string) error
}
