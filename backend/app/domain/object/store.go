package object

import (
	"time"
)

type Store struct {
	ID        int          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string       `gorm:"size:255;not null" json:"name"`
	ImageURL  *string      `json:"image_url"`
	EventID   int          `gorm:"not null" json:"event_id"`
	Event     Event        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Staffs    []StoreStaff `gorm:"foreignKey:StoreID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Items     []Item       `gorm:"foreignKey:StoreID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
