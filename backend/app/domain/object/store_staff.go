package object

import (
	"time"

)

type StoreStaff struct {
	ID            int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Role          int       `gorm:"not null" json:"role"`
        StudentID     int       `gorm:"not null" json:"student_id"`
	Student       Student   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	StoreID       int       `gorm:"not null" json:"store_id"`
	Store         Store     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Orders        []Order   `gorm:"foreignKey:StoreStaffID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func NewStoreStaff(studentId int, role int, storeId int) (*StoreStaff, error) {
	storeStaff := &StoreStaff{
		StudentID:     studentId,
		Role:          role,
		StoreID:       storeId,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	return storeStaff, nil
}
