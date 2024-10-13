package order

import (
	"backend/app/usecase"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Group, u usecase.Order) {
	h := &orderHandler{
		orderUseCase: u,
	}

	e.POST("/orders", h.Create)
}
