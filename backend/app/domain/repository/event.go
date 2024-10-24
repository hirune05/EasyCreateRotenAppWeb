package repository

import (
	"context"

	"backend/app/domain/object"

	"gorm.io/gorm"
)

type EventRepository interface {
	Create(ctx context.Context, tx *gorm.DB, event *object.Event) error
	GetByID(ctx context.Context, id string) (*object.Event, error)
	GetAll(ctx context.Context) ([]*object.Event, error)
	GetAllNowOn(ctx context.Context) ([]*object.Event, error)
	Delete(ctx context.Context, id string) error
}
