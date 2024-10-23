package student

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	StudentID     int    `json:"student_id"`
	Password      string `json:"password"`
        EventID       int    `json:"event_id"`
}

type LoginResponse struct {
	Token         string `json:"token"`
	StudentID     int    `json:"student_id"`
	Name          string `json:"name"`
}

func (h *studentHandler) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	ctx := c.Request().Context()

	dto, err := h.studentUseCase.Login(ctx, req.StudentID, req.Password, req.EventID)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}

	return c.JSON(http.StatusOK, LoginResponse{
		Token:         dto.Token,
		StudentID:     dto.StudentID,
		Name:          dto.Name,
	})
}
