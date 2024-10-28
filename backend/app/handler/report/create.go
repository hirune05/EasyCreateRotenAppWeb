package report

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AddRequest struct {
        StoreID       *int    `json:"storeId"`
        StoreStaffID  *int    `json:"storeStaffId"`
        Title         string  `json:"title"`
        Description   string  `json:"description"`
}

func (h *reportHandler) Create(c echo.Context) error {
	var req AddRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	ctx := c.Request().Context()

	dto, err := h.reportUseCase.Create(ctx, req.StoreID, req.StoreStaffID, req.Title, req.Description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, dto.Report)
}
