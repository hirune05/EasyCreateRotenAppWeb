package order

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *orderHandler) GetByStoreID(c echo.Context) error {
	id := c.Param("storeId")

	dto, err := h.orderUseCase.GetByStoreID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "order not found"})
	}

	return c.JSON(http.StatusOK, dto.Orders)
}
