package adminUser

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Request body for `POST /v1/adminUsers`
type AddRequest struct {
	Username string
	Email    string
	Password string
}

func (h *adminUserHandler) Create(c echo.Context) error {
	var req AddRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	ctx := c.Request().Context()

	dto, err := h.adminUserUseCase.Create(ctx, req.Username, req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, dto.AdminUser)
}
