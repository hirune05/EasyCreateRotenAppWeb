package student

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	StudentId int    `json:"student_id"`
	Password      string `json:"password"`
}

type LoginResponse struct {
	Token         string `json:"token"`
	StudentId     int    `json:"student_id"`
	Name          string `json:"name"`
}

func (h *studentHandler) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	ctx := c.Request().Context()

	dto, err := h.studentUseCase.Login(ctx, req.StudentId, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}

	return c.JSON(http.StatusOK, LoginResponse{
		Token:         dto.Token,
		StudentId:     dto.StudentId,
		Name:          dto.Name,
	})
}
