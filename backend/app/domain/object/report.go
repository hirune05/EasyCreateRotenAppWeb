package object

import (
	"time"
)

type Report struct {
	ID           int        `db:"id" gorm:"primaryKey;autoIncrement" json:"id"`
	StoreID      *int       `db:"store_id" json:"storeId"`
	Store        *Store      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"Store"`
	StoreStaffID *int       `db:"store_staff_id" json:"storeStaffId"`
	StoreStaff   *StoreStaff `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"StoreStaff"`
	Description  string     `db:"description" gorm:"not null" json:"description"`
	CreatedAt    time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt    time.Time  `db:"updated_at" json:"updatedAt"`
}
