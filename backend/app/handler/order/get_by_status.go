package order

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *orderHandler) GetByStatus(c echo.Context) error {
	storeId := c.Param("storeId")
	status := c.Param("status")

	order, err := h.orderUseCase.GetByStatus(c.Request().Context(), storeId, status)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "order not found"})
	}

	return c.JSON(http.StatusOK, order)
}
