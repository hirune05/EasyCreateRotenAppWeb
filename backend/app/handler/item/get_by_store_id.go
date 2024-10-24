package item

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *itemHandler) GetByStoreID(c echo.Context) error {
	id := c.Param("storeId")

	item, err := h.itemUseCase.GetByStoreID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "item not found"})
	}

	return c.JSON(http.StatusOK, item)
}
