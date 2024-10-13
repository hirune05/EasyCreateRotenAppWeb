// @Summary Get an order by ID
// @Description Get the details of a specific order by its ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Success 200 {object} object.Order
// @Failure 404 {object} map[string]string "Order not found"
// @Router /v1/orders/{id} [get]
package order

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *orderHandler) GetByID(c echo.Context) error {
	id := c.Param("id")

	order, err := h.orderUseCase.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "order not found"})
	}

	return c.JSON(http.StatusOK, order)
}
