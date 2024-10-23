// @Summary Get an event by ID
// @Description Get the details of a specific event by its ID
// @Tags events
// @Accept  json
// @Produce  json
// @Param id path int true "Event ID"
// @Success 200 {object} object.Event
// @Failure 404 {object} map[string]string "Event not found"
// @Router /v1/events/{id} [get]
package event

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *eventHandler) GetByID(c echo.Context) error {
	id := c.Param("id")

	event, err := h.eventUseCase.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "event not found"})
	}

	return c.JSON(http.StatusOK, event)
}
