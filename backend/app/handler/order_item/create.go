package orderItem

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AddRequest struct {
	OrderID  int
	ItemID   int
	Quantity int
	Arranges *string
}

func (h *orderItemHandler) CreateOrderItem(c echo.Context) error {
	var req AddRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	ctx := c.Request().Context()

	dto, err := h.orderItemUseCase.Create(ctx, req.OrderID, req.ItemID, req.Quantity, req.Arranges)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, dto.OrderItem)
}
