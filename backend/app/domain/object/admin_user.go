package object

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AdminUser struct {
	ID        int       `db:"id" gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `db:"username" gorm:"size:255;not null" json:"username"`
	Password  string    `db:"password" gorm:"size:255;not null" json:"password"`
	Email     string    `db:"email" gorm:"size:255;not null;unique" json:"email"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
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

func (a *AdminUser) CheckPassword(pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(pass)) == nil
}

func (a *AdminUser) SetPassword(pass string) error {
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
