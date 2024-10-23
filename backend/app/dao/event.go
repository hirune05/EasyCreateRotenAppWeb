package dao

import (
	"backend/app/domain/object"
	"backend/app/domain/repository"
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

var _ repository.EventRepository = (*EventRepositoryImpl)(nil)

type EventRepositoryImpl struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepositoryImpl {
	return &EventRepositoryImpl{db: db}
}

func (r *EventRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, event *object.Event) error {
	if err := tx.WithContext(ctx).Create(event).Error; err != nil {
		return fmt.Errorf("failed to insert event: %w", err)
	}
	return nil
}

func (r *EventRepositoryImpl) GetByID(ctx context.Context, id string) (*object.Event, error) {
	var event object.Event

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&event).Error; err != nil {
		return nil, fmt.Errorf("failed to find event by id: %w", err)
	}

	return &event, nil
}

func (r *EventRepositoryImpl) GetAll(ctx context.Context) ([]*object.Event, error) {
	var events []*object.Event

	if err := r.db.WithContext(ctx).Find(&events).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve all events: %w", err)
	}

	return events, nil
}

func (r *EventRepositoryImpl) GetAllNowOn(ctx context.Context) ([]*object.Event, error) {
	var events []*object.Event

	if err := r.db.WithContext(ctx).Where("? BETWEEN start_time AND end_time", time.Now()).Find(&events).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve all events: %w", err)
	}

	return events, nil
}

func (r *EventRepositoryImpl) Delete(ctx context.Context, id string) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&object.Event{}).Error; err != nil {
		return fmt.Errorf("failed to delete event: %w", err)
	}
	return nil
}
