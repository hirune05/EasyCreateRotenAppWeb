package object

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type StoreStaff struct {
	ID            int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string    `gorm:"size:255;not null" json:"name"`
	Password      string    `gorm:"size:255;not null" json:"password"`
	StudentNumber int       `gorm:"not null" json:"student_number"`
	Role          int       `gorm:"not null" json:"role"`
	StoreID       *int      `json:"store_id"`
	Store         Store     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Orders        []Order   `gorm:"foreignKey:StoreStaffID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func NewStoreStaff(name, password string, studentNumber, role int, storeId *int) (*StoreStaff, error) {
	adminUser := &StoreStaff{
		Name:          name,
		Password:      password,
		StudentNumber: studentNumber,
		Role:          role,
		StoreID:       storeId,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := adminUser.SetPassword(password); err != nil {
		return nil, fmt.Errorf("パスワード設定エラー: %w", err)
	}

	return adminUser, nil
}

// 指定されたパスワードが adminUser のパスワードと一致するかを確認
func (a *StoreStaff) CheckPassword(pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(pass)) == nil
}

// パスワードをハッシュ化し、adminUser に設定する
func (a *StoreStaff) SetPassword(pass string) error {
	// パスワードが空でないか確認
	if pass == "" {
		return fmt.Errorf("パスワードは空にできません")
	}

	passwordHash, err := generatePasswordHash(pass)
	if err != nil {
		return fmt.Errorf("パスワードハッシュ生成エラー: %w", err)
	}
	a.Password = passwordHash
	return nil
}

func generatePasswordHash(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("hashing password failed: %w", err)
	}
	return string(hash), nil
}
