package order

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"backend/app/usecase"
)

type AddComplexRequest struct {
	StoreID      int                          `json:"store_id"`
	StoreStaffID int                          `json:"store_staff_id"`
        Items        []*usecase.OrderComplexItem  `json:"items"`
}

func (h *orderHandler) CreateComplex(c echo.Context) error {
	var req AddComplexRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	ctx := c.Request().Context()

	dto, err := h.orderUseCase.CreateComplex(ctx, req.StoreID, req.StoreStaffID, req.Items)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, dto.Order)
}
