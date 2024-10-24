package item

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *itemHandler) GetAll(c echo.Context) error {
	items, err := h.itemUseCase.GetAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to retrieve items"})
	}

	return c.JSON(http.StatusOK, items)
}
