package object

import (
	"time"
)

type StoreStaff struct {
	ID            int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string    `gorm:"size:255;not null" json:"name"`
	Password      string    `gorm:"size:255;not null" json:"password"`
	StudentNumber int       `gorm:"not null" json:"student_number"`
	Role          int       `gorm:"not null" json:"role"`
	StoreID       int       `gorm:"not null" json:"store_id"`
	Store         Store     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Orders        []Order   `gorm:"foreignKey:StoreStaffID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
