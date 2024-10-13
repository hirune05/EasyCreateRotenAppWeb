package object

import (
	"time"
)

type Report struct {
	ID           int        `gorm:"primaryKey;autoIncrement" json:"id"`
	StoreID      *int       `json:"store_id"`
	Store        Store      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	StoreStaffID *int       `json:"store_staff_id"`
	StoreStaff   StoreStaff `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Description  string     `gorm:"not null" json:"description"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
