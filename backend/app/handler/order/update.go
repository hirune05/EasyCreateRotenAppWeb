package order

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type UpdateRequest struct {
	StoreID      int
	StoreStaffID int
	Status       int
	PickedUpAt   *time.Time
}

func (h *orderHandler) Update(c echo.Context) error {
	id := c.Param("id")
	var req UpdateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	order, err := h.orderUseCase.Update(c.Request().Context(), id, req.StoreID, req.StoreStaffID, req.Status, req.PickedUpAt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to update order"})
	}

	return c.JSON(http.StatusOK, order)
}
