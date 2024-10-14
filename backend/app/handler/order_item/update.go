package orderItem

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UpdateRequest struct {
	Quantity int
	Arranges *string
}

func (h *orderItemHandler) Update(c echo.Context) error {
	id := c.Param("id")
	var req UpdateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	order, err := h.orderItemUseCase.Update(c.Request().Context(), id, req.Quantity, req.Arranges)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update order"})
	}

	return c.JSON(http.StatusOK, order)
}
