package object

import (
	"time"
)

type Item struct {
	ID          int         `db:"id" gorm:"primaryKey;autoIncrement" json:"id"`
	StoreID     int         `db:"store_id" gorm:"not null" json:"storeId"`
	Store       *Store      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"Store"`
	Name        string      `db:"name" gorm:"size:255;not null" json:"name"`
	Description *string     `db:"description" json:"description"`
	Price       int         `db:"price" gorm:"not null" json:"price"`
	ImageURL    *string     `db:"image_url" json:"imageUrl"`
	CreatedAt   time.Time   `db:"created_at" json:"createdAt"`
	UpdatedAt   time.Time   `db:"updated_at" json:"updatedAt"`
	OrderItems  []*OrderItem `gorm:"foreignKey:ItemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"OrderItems"`
}

func NewItem(storeID int, name string, description *string, price int, imageURL *string) (*Item, error) {
        item := &Item{
		StoreID:      storeID,
                Name:         name,
                Description:  description,
                Price:        price,
                ImageURL:     imageURL,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
        }
        return item, nil
}
