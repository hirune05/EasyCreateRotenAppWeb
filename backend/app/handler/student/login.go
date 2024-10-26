package student

import (
	"net/http"
        "log"

	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	StudentID     int    `json:"studentId"`
	Password      string `json:"password"`
        EventID       int    `json:"eventId"`
}

type LoginResponse struct {
	Token         string `json:"token"`
	StudentID     int    `json:"studentId"`
	Name          string `json:"name"`
        StoreID       int    `json:"storeId"`
        StoreName     string `json:"storeName"`
}

func (h *studentHandler) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	ctx := c.Request().Context()

	dto, err := h.studentUseCase.Login(ctx, req.StudentID, req.Password, req.EventID)
	if err != nil {
                log.Printf("Request error: %v", err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}

	return c.JSON(http.StatusOK, LoginResponse{
		Token:         dto.Token,
		StudentID:     dto.StudentID,
		Name:          dto.Name,
                StoreID:       dto.StoreID,
                StoreName:     dto.StoreName,
	})
}
