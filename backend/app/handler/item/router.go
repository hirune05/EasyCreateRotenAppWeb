package item

import (
	"backend/app/usecase"

	"github.com/labstack/echo/v4"
)

type itemHandler struct {
	itemUseCase usecase.Item
}

func RegisterRoutes(e *echo.Group, u usecase.Item) {
	h := &itemHandler{
		itemUseCase: u,
	}

	e.POST("/items", h.Create)
	e.GET("/items", h.GetAll)
	e.GET("/items/:id", h.GetByID)
	e.GET("/items/store/:storeId", h.GetByStoreID)
}
