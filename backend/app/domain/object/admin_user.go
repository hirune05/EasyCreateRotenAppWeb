package object

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AdminUser struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"size:255;not null" json:"username"`
	Password  string    `gorm:"size:255;not null" json:"password"`
	Email     string    `gorm:"size:255;not null" json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewAdminUser(username, email, password string) (*AdminUser, error) {
	adminUser := &AdminUser{
		Username:  username,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := adminUser.SetPassword(password); err != nil {
		return nil, fmt.Errorf("パスワード設定エラー: %w", err)
	}

	return adminUser, nil
}

// 指定されたパスワードが adminUser のパスワードと一致するかを確認
func (a *AdminUser) CheckPassword(pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(pass)) == nil
}

// パスワードをハッシュ化し、adminUser に設定する
func (a *AdminUser) SetPassword(pass string) error {
	// パスワードが空でないか確認
	if pass == "" {
		return fmt.Errorf("パスワードは空にできません")
	}

	passwordHash, err := generateAdminPasswordHash(pass)
	if err != nil {
		return fmt.Errorf("パスワードハッシュ生成エラー: %w", err)
	}
	a.Password = passwordHash
	return nil
}

func generateAdminPasswordHash(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("hashing password failed: %w", err)
	}
	return string(hash), nil
}
