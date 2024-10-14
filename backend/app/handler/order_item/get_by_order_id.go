package orderItem

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *orderItemHandler) GetByOrderId(c echo.Context) error {
	id := c.Param("id")

	orderItems, err := h.orderItemUseCase.GetByOrderId(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to retrieve order item"})
	}

	if orderItems == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "order item not found"})
	}

	return c.JSON(http.StatusOK, orderItems)
}
