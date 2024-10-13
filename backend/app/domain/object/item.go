package object

import (
	"time"
)

type Item struct {
	ID          int         `gorm:"primaryKey;autoIncrement" json:"id"`
	StoreID     int         `gorm:"not null" json:"store_id"`
	Store       Store       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Name        string      `gorm:"size:255;not null" json:"name"`
	Description *string     `json:"description"`
	Price       int         `gorm:"not null" json:"price"`
	ImageURL    *string     `json:"image_url"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	OrderItems  []OrderItem `gorm:"foreignKey:ItemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // 関連するOrderItems
}
