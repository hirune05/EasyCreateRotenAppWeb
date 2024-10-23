package student

import (
	"backend/app/usecase"

	"github.com/labstack/echo/v4"
)

type studentHandler struct {
	studentUseCase usecase.Student
}

func RegisterRoutes(e *echo.Group, u usecase.Student) {
	h := &studentHandler{
		studentUseCase: u,
	}

	e.POST("/student", h.Create)
	e.POST("/student/login", h.Login)
}
