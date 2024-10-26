package storeStaff

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AddRequest struct {
	StudentID     int   `json:"studentId"`
	Role          int   `json:"role"`
	StoreID       int   `json:"storeId"`
}

func (h *storeStaffHandler) Create(c echo.Context) error {
	var req AddRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	ctx := c.Request().Context()

	dto, err := h.storeStaffUseCase.Create(ctx, req.StudentID, req.Role, req.StoreID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, dto.StoreStaff)
}
