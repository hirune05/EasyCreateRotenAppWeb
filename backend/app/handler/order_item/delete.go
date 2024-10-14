package orderItem

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *orderItemHandler) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := h.orderItemUseCase.Delete(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete order_item"})
	}

	return c.NoContent(http.StatusNoContent)
}
