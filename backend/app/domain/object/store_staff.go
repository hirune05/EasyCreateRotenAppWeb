package object

import (
	"time"

)

type StoreStaff struct {
	ID            int       `db:"id" gorm:"primaryKey;autoIncrement" json:"id"`
	Role          int       `db:"role" gorm:"not null" json:"role"`
        StudentID     int       `db:"student_id" gorm:"not null;uniqueIndex:idx_student_store" json:"studentId"`
	Student       *Student   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"Student"`
	StoreID       int       `db:"store_id" gorm:"not null" json:"storeId;uniqueIndex:idx_student_store"`
	Store         *Store     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"Store"`
	CreatedAt     time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt     time.Time `db:"updated_at" json:"updatedAt"`
	Orders        []*Order   `gorm:"foreignKey:StoreStaffID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"Orders"`
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
