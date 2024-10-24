package order

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type SetPickedRequest struct {
        PickedUpAt   time.Time  `json:"picked_up_at"`
}

func (h *orderHandler) SetPickedUpAt(c echo.Context) error {
	id := c.Param("id")
	var req SetPickedRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	ctx := c.Request().Context()

	dto, err := h.orderUseCase.Update(ctx, id, nil, nil, nil, &req.PickedUpAt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, dto.Order)
}
