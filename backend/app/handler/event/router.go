package event

import (
	"backend/app/usecase"

	"github.com/labstack/echo/v4"
)

type eventHandler struct {
	eventUseCase usecase.Event
}

func RegisterRoutes(e *echo.Group, u usecase.Event) {
	h := &eventHandler{
		eventUseCase: u,
	}

	e.GET("/events", h.GetAll)
	e.GET("/events/now", h.GetAllNowOn)
	e.GET("/events/:id", h.GetByID)
}
