package student

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AddRequest struct {
	Id            int     `json:"id"`
	Name          string  `json:"name"`
	Password      string  `json:"password"`
}

func (h *studentHandler) Create(c echo.Context) error {
	var req AddRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	ctx := c.Request().Context()

	dto, err := h.studentUseCase.Create(ctx, req.Id, req.Name, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, dto.Student)
}
