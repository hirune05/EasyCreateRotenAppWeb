package storeStaff

import (
	"backend/app/usecase"

	"github.com/labstack/echo/v4"
)

type storeStaffHandler struct {
	storeStaffUseCase usecase.StoreStaff
}

func RegisterRoutes(e *echo.Group, u usecase.StoreStaff) {
	h := &storeStaffHandler{
		storeStaffUseCase: u,
	}

	e.POST("/storeStaff", h.Create)
	e.POST("/storeStaff/login", h.Login)
}
