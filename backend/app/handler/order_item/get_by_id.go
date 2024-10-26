package orderItem

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *orderItemHandler) GetByID(c echo.Context) error {
	id := c.Param("id")

	dto, err := h.orderItemUseCase.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "order item not found"})
	}

	return c.JSON(http.StatusOK, dto.OrderItem)
}
