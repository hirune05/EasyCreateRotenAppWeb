package order

import (
	"backend/app/usecase"

	"github.com/labstack/echo/v4"
)

type orderHandler struct {
	orderUseCase usecase.Order
}

func RegisterRoutes(e *echo.Group, u usecase.Order) {
	h := &orderHandler{
		orderUseCase: u,
	}

	e.POST("/orders", h.Create)
	e.GET("/orders", h.GetAll)
	e.POST("/orders/complex", h.CreateComplex)
	e.GET("/orders/:id", h.GetByID)
	e.GET("/orders/store/:storeId", h.GetByStoreID)
	e.GET("/orders/store/:storeId/:status", h.GetByStatus)
	e.PUT("/orders/:id", h.Update)
        e.PUT("/orders/:id/picked", h.SetPickedUpAt)
        e.PUT("/orders/:id/status", h.SetStatus)
	e.DELETE("/orders/:id", h.Delete)
}
