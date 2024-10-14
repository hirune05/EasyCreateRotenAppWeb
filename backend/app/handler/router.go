package handler

import (
	"net/http"
	"time"

	"backend/app/handler/order"
	orderItem "backend/app/handler/order_item"
	"backend/app/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(ou usecase.Order, oiu usecase.OrderItem) http.Handler {
	e := echo.New()

	// A good base middleware stack
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		//TODO allow frontend IP
	}))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 60 * time.Second,
	}))

	v1 := e.Group("/v1")

	order.RegisterRoutes(v1, ou)
	orderItem.RegisterRoutes(v1, oiu)

	return e
}
