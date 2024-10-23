package object

import (
	"time"
)

type Event struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:255;not null" json:"name"`
	Year      int       `gorm:"not null" json:"year"`
        StartTime time.Time `gorm:"not null" json:"start_time"`
        EndTime   time.Time `gorm:"not null" json:"end_time"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewEvent(name string, year int, startTime time.Time, endTime time.Time) (*Event, error) {
	order := &Event{
                Name:         name,
                Year:         year,
                StartTime:    startTime,
                EndTime:      endTime,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return order, nil
}
