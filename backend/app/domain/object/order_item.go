package object

import (
	"time"
)

type OrderItem struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID   int       `gorm:"not null" json:"order_id"`
	Order     Order     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ItemID    int       `gorm:"not null" json:"item_id"`
	Item      Item      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Arranges  *string   `json:"arranges"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewOrderItem(orderID, itemID, quantity int, arranges *string) (*OrderItem, error) {
	orderItem := &OrderItem{
		OrderID:   orderID,
		ItemID:    itemID,
		Quantity:  quantity,
		Arranges:  arranges,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return orderItem, nil
}
