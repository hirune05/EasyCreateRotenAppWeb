package object

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Student struct {
	ID int       `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:255;not null" json:"username"`
	Password  string    `gorm:"size:255;not null" json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

func NewStudent(id int, name string, password string) (*Student, error) {
	student := &Student{
                ID:  id,
		Name:       name,
                Password:   password,
	}

	if err := student.SetPassword(password); err != nil {
		return nil, fmt.Errorf("パスワード設定エラー: %w", err)
	}

	return student, nil
}

func (a *Student) CheckPassword(pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(pass)) == nil
}

func (a *Student) SetPassword(pass string) error {
	if pass == "" {
		return fmt.Errorf("パスワードは空にできません")
	}

	passwordHash, err := generateStudentPasswordHash(pass)
	if err != nil {
		return fmt.Errorf("パスワードハッシュ生成エラー: %w", err)
	}
	a.Password = passwordHash
	return nil
}

func generateStudentPasswordHash(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("hashing password failed: %w", err)
	}
	return string(hash), nil
}

func (a *Student) GenerateJWTToken(secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":             a.ID,
		"exp":            time.Now().Add(time.Hour * 6).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
