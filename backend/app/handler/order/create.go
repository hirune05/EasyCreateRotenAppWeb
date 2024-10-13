package order

import (
	"backend/app/usecase"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type orderHandler struct {
	orderUseCase usecase.Order
}

// Request body for `POST /v1/orders`
type AddRequest struct {
	StoreID      int
	StoreStaffID int
	Status       int
	PickedUpAt   *time.Time
}

func (h *orderHandler) Create(c echo.Context) error {
	var req AddRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	ctx := c.Request().Context()

	dto, err := h.orderUseCase.Create(ctx, req.StoreID, req.StoreStaffID, req.Status, req.PickedUpAt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, dto.Order)
}
