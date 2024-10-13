package object

import (
	"time"
)

type Order struct {
	ID           int         `gorm:"primaryKey;autoIncrement" json:"id"`
	StoreID      int         `gorm:"not null" json:"store_id"`
	Store        Store       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	PickedUpAt   time.Time   `json:"picked_up_at"`
	Status       int         `gorm:"not null" json:"status"`
	StoreStaffID int         `gorm:"not null" json:"store_staff_id"`
	StoreStaff   StoreStaff  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` 
}
