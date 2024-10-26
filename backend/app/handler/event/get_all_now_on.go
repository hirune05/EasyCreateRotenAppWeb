package event

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *eventHandler) GetAllNowOn(c echo.Context) error {
	dto, err := h.eventUseCase.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to retrieve events"})
	}

	return c.JSON(http.StatusOK, dto.Events)
}
