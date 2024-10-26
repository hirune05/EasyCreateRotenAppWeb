package orderItem

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *orderItemHandler) GetByOrderID(c echo.Context) error {
	id := c.Param("id")

	dto, err := h.orderItemUseCase.GetByOrderID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to retrieve order item"})
	}

	return c.JSON(http.StatusOK, dto.OrderItems)
}
