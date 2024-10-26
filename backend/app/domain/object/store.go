package object

import (
	"time"
)

type Store struct {
	ID        int          `db:"id" gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string       `db:"name" gorm:"size:255;not null" json:"name"`
	ImageURL  *string      `db:"image_url" json:"imageUrl"`
	EventID   int          `db:"event_id" gorm:"not null" json:"eventId"`
	Event     *Event        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"Event"`
	CreatedAt time.Time    `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time    `db:"updated_at" json:"updatedAt"`
	Staffs    []*StoreStaff `gorm:"foreignKey:StoreID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:Staffs`
	Items     []*Item       `gorm:"foreignKey:StoreID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:Items`
}
