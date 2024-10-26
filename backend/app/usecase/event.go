package usecase

import (
	"backend/app/domain/object"
	"backend/app/domain/repository"
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Event interface {
	Create (ctx context.Context, name string, year int, startTime time.Time, endTime time.Time) (*EventDTO, error)
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*EventDTO, error)
	GetAll(ctx context.Context) (*EventsDTO, error)
	GetAllNowOn(ctx context.Context) (*EventsDTO, error)
}
type event struct {
	db        *gorm.DB
	eventRepo repository.EventRepository
}

var _ Event = (*event)(nil)

type EventDTO struct {
        Event *object.Event
}
type EventsDTO struct {
        Events []*object.Event
}


func NewEvent(db *gorm.DB, eventRepo repository.EventRepository) *event {
	return &event{
		db:        db,
		eventRepo: eventRepo,
	}
}

func (a *event) Create(ctx context.Context, name string, year int, startTime time.Time, endTime time.Time) (*EventDTO, error) {
	acc, err := object.NewEvent(name, year, startTime, endTime)
	if err != nil {
		return nil, err
	}

	tx := a.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err := a.eventRepo.Create(ctx, tx, acc); err != nil {
		return nil, err
	}

	return &EventDTO{
		Event: acc,
	}, nil
}

func (a *event) Delete(ctx context.Context, id string) error {
	tx := a.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err := a.eventRepo.Delete(ctx, id); err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete event: %w", err)
	}

	return nil
}

func (a *event) GetByID(ctx context.Context, id string) (*EventDTO, error) {
	event, err := a.eventRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get event by id: %w", err)
	}

	return &EventDTO{
		Event: event,
	}, nil
}

func (r *event) GetAll(ctx context.Context) (*EventsDTO, error) {
	events, err := r.eventRepo.GetAll(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve all events: %w", err)
	}

	return &EventsDTO{
		Events: events,
	}, nil
}

func (r *event) GetAllNowOn(ctx context.Context) (*EventsDTO, error) {
	events, err := r.eventRepo.GetAllNowOn(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve all events: %w", err)
	}

	return &EventsDTO{
		Events: events,
	}, nil
}
