package object

import (
	"time"
)

type Event struct {
	ID        int       `db:"id" gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `db:"name" gorm:"size:255;not null;unique" json:"name"`
	Year      int       `db:"year" gorm:"not null" json:"year"`
        StartTime time.Time `db:"start_time" gorm:"not null" json:"startTime"`
        EndTime   time.Time `db:"end_time" gorm:"not null" json:"endTime"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
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
