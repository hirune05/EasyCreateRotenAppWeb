package storeStaff

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	StudentNumber int    `json:"student_number"`
	Password      string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (h *storeStaffHandler) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	ctx := c.Request().Context()

	dto, err := h.storeStaffUseCase.Login(ctx, req.StudentNumber, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}

	return c.JSON(http.StatusOK, LoginResponse{Token: dto.Token})
}
