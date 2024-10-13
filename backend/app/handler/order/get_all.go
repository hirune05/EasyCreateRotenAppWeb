package order

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *orderHandler) GetAll(c echo.Context) error {
	orders, err := h.orderUseCase.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to retrieve orders"})
	}

	return c.JSON(http.StatusOK, orders)
}
