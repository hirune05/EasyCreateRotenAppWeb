package item

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *itemHandler) GetByID(c echo.Context) error {
	id := c.Param("id")

	item, err := h.itemUseCase.GetByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "item not found"})
	}

	return c.JSON(http.StatusOK, item)
}
