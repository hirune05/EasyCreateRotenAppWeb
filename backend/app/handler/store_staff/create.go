package storeStaff

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AddRequest struct {
	Name          string
	Password      string
	StudentNumber int
	Role          int
	StoreId       *int
}

func (h *storeStaffHandler) Create(c echo.Context) error {
	var req AddRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	ctx := c.Request().Context()

	dto, err := h.storeStaffUseCase.Create(ctx, req.Name, req.Password, req.StudentNumber, req.Role, req.StoreId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, dto.StoreStaff)
}
