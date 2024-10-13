package repository

import (
	"context"

	"backend/app/domain/object"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(ctx context.Context, tx *gorm.DB, acc *object.Order) error
}
