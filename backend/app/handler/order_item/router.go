package orderItem

import (
	"backend/app/usecase"

	"github.com/labstack/echo/v4"
)

type orderItemHandler struct {
	orderItemUseCase usecase.OrderItem
}

func RegisterRoutes(e *echo.Group, u usecase.OrderItem) {
	h := &orderItemHandler{
		orderItemUseCase: u,
	}

	e.POST("/order-items", h.CreateOrderItem)
	e.GET("/order-items/:id", h.GetByOrderId)

}
