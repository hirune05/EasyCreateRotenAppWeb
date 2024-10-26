package item

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AddRequest struct {
        StoreID       int     `json:"storeId"`
        Name          string  `json:"name"`
        Description   *string `json:"description"`
        Price         int     `json:"price"`
        ImageURL      *string `json:"imageUrl"`
}

func (h *itemHandler) Create(c echo.Context) error {
	var req AddRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	ctx := c.Request().Context()

	dto, err := h.itemUseCase.Create(ctx, req.StoreID, req.Name, req.Description, req.Price, req.ImageURL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, dto.Item)
}
