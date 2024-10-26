package object

import (
	"time"
)

type OrderItem struct {
	ID        int       `db:"id" gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID   int       `db:"order_id" gorm:"not null" json:"orderId"`
	Order     *Order     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:Order`
	ItemID    int       `db:"item_id" gorm:"not null" json:"itemId"`
	Item      *Item      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:Item`
	Arranges  *string   `db:"arranges" json:"arranges"`
	Quantity  int       `db:"quantity" gorm:"not null" json:"quantity"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
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
