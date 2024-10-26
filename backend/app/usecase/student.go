package usecase

import (
	"backend/app/domain/object"
	"backend/app/domain/repository"
	"context"
	"errors"
	"os"

	"gorm.io/gorm"
)

type Student interface {
	Create(ctx context.Context, id int, name string, password string) (*StudentDTO, error)
	Login(ctx context.Context, studentNumber int, password string, EventID int) (*LoginStudentDTO, error)
}
type LoginStudentDTO struct {
	Token        string `json:"token"`
	StudentID    int    `json:"studentId"`
	Name         string `json:"name"`
	StoreID      int    `json:"storeId"`
	StoreName    string `json:"storeName"`
	StoreStaffID int    `json:"storeStaffId"`
}

type StudentDTO struct {
	Student *object.Student
}

type student struct {
	db          *gorm.DB
	studentRepo repository.StudentRepository
}

var _ Student = (*student)(nil)

func NewStudent(db *gorm.DB, studentRepo repository.StudentRepository) *student {
	return &student{
		db:          db,
		studentRepo: studentRepo,
	}
}

func (a *student) Create(ctx context.Context, id int, name string, password string) (*StudentDTO, error) {
	student, err := object.NewStudent(id, name, password)
	if err != nil {
		return nil, err
	}

	tx := a.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	if err := a.studentRepo.Create(ctx, tx, student); err != nil {
		return nil, err
	}

	return &StudentDTO{
		Student: student,
	}, nil
}

func (a *student) Login(ctx context.Context, studentID int, password string, eventID int) (*LoginStudentDTO, error) {
	student, err := a.studentRepo.FindByStudentId(ctx, studentID)
	if err != nil {
		return nil, err
	}

	if !student.CheckPassword(password) {
		return nil, errors.New("invalid student id or password")
	}

	dao, err := a.studentRepo.CheckInEvent(ctx, eventID, studentID)
	if err != nil {
		return nil, err
	}
	if dao.StoreID == 0 && dao.StoreName == "" && dao.StoreStaffID == 0 {
		return nil, errors.New("student id is not a vendor at this event.")
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return nil, errors.New("JWT secret key is not configured")
	}

	token, err := student.GenerateJWTToken(secretKey)
	if err != nil {
		return nil, err
	}
	return &LoginStudentDTO{
		Token:        token,
		StudentID:    student.ID,
		Name:         student.Name,
		StoreID:      dao.StoreID,
		StoreName:    dao.StoreName,
		StoreStaffID: dao.StoreStaffID,
	}, nil
}
