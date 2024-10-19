package object

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type StoreStaff struct {
	ID            int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string    `gorm:"size:255;not null" json:"name"`
	Password      string    `gorm:"size:255;not null" json:"password"`
	StudentNumber int       `gorm:"not null;unique" json:"student_number"`
	Role          int       `gorm:"not null" json:"role"`
	StoreID       int       `gorm:"not null" json:"store_id"`
	Store         Store     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Orders        []Order   `gorm:"foreignKey:StoreStaffID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func NewStoreStaff(name, password string, studentNumber, role int, storeId int) (*StoreStaff, error) {
	storeStaff := &StoreStaff{
		Name:          name,
		Password:      password,
		StudentNumber: studentNumber,
		Role:          role,
		StoreID:       storeId,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := storeStaff.SetPassword(password); err != nil {
		return nil, fmt.Errorf("パスワード設定エラー: %w", err)
	}

	return storeStaff, nil
}

func (a *StoreStaff) CheckPassword(pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(pass)) == nil
}

func (a *StoreStaff) SetPassword(pass string) error {
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

func (a *StoreStaff) GenerateJWTToken(secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":             a.ID,
		"student_number": a.StudentNumber,
		"exp":            time.Now().Add(time.Hour * 72).Unix(), // Token expiration time
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
