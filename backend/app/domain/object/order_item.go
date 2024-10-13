package object

import (
	"time"
)

type OrderItem struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID   int       `gorm:"not null" json:"order_id"`                      // 外部キー
	Order     Order     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Orderとのリレーション
	ItemID    int       `gorm:"not null" json:"item_id"`                       // 外部キー
	Item      Item      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // Itemとのリレーション
	Arranges  *string   `json:"arranges"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
