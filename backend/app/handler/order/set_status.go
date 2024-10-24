package order

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type SetStatusRequest struct {
        Status       int  `json:"status"`
}

func (h *orderHandler) SetStatus(c echo.Context) error {
	id := c.Param("id")
	var req SetStatusRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	ctx := c.Request().Context()

	dto, err := h.orderUseCase.Update(ctx, id, nil, nil, &req.Status, nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, dto.Order)
}
