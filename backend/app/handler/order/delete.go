package order

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *orderHandler) Delete(c echo.Context) error {
	id := c.Param("id")

	if err := h.orderUseCase.Delete(c.Request().Context(), id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to delete order"})
	}

	return c.NoContent(http.StatusNoContent)
}
