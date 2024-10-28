package report

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *reportHandler) SendEmail(c echo.Context) error {

	dto, err := h.reportUseCase.SendEmail(c.Request().Context(), "Title", "Body")
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "send email failed + " + err.Error()})
	}

	return c.JSON(http.StatusOK, dto.Message)
}
