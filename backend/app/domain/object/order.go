package object

import (
	"time"
)

type Order struct {
	ID           int         `db:"id" gorm:"primaryKey;autoIncrement" json:"id"`
	StoreID      int         `db:"store_id" gorm:"not null" json:"storeId"`
	Store        *Store       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"Store"`
	PickedUpAt   *time.Time  `db:"picked_up_at" json:"pickedUpAt"`
	Status       int         `db:"status" gorm:"not null" json:"status"`
	StoreStaffID int         `db:"store_staff_id" gorm:"not null" json:"storeStaffId"`
	StoreStaff   *StoreStaff  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"StoreStaff"`
	CreatedAt    time.Time   `db:"created_at" json:"createdAt"`
	UpdatedAt    time.Time   `db:"updated_at" json:"updatedAt"`
	OrderItems   []*OrderItem `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"OrderItems"`
}

func NewOrder(storeID, storeStaffID, status int, pickedUpAt *time.Time) (*Order, error) {
	order := &Order{
		StoreID:      storeID,
		Status:       status,
		StoreStaffID: storeStaffID,
		PickedUpAt:   pickedUpAt,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return order, nil
}
