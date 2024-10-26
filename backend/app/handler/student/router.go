package student

import (
	"backend/app/usecase"
	"backend/app/appmiddleware"

	"github.com/labstack/echo/v4"
)

type studentHandler struct {
	studentUseCase usecase.Student
}

func RegisterRoutes(e *echo.Group, u usecase.Student) {
	h := &studentHandler{
		studentUseCase: u,
	}

	e.POST("/student", h.Create, appmiddleware.RequestAuthHandker)
	e.POST("/student/login", h.Login)
}
