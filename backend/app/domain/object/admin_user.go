package object

import (
	"time"
)

type AdminUser struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"size:255;not null" json:"username"`
	Password  string    `gorm:"size:255;not null" json:"password"`
	Email     string    `gorm:"size:255;not null" json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
