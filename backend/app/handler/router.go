package handler

import (
	"fmt"
	"net/http"
	"os"
	"time"

	adminUser "backend/app/handler/admin_user"
	event "backend/app/handler/event"
	item "backend/app/handler/item"
	order "backend/app/handler/order"
	orderItem "backend/app/handler/order_item"
	report "backend/app/handler/report"
	storeStaff "backend/app/handler/store_staff"
	student "backend/app/handler/student"
	"backend/app/usecase"

	"backend/app/appmiddleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(ou usecase.Order, oiu usecase.OrderItem, au usecase.AdminUser, su usecase.StoreStaff, stu usecase.Student, eu usecase.Event, iu usecase.Item, ru usecase.Report) http.Handler {
	e := echo.New()

	// A good base middleware stack
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL")},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowCredentials: true,
	}))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 60 * time.Second,
	}))

	v1_noToken := e.Group("/v1")
	student.RegisterRoutes(v1_noToken, stu)
	event.RegisterRoutes(v1_noToken, eu)

	v1 := e.Group("/v1")
	v1.Use(appmiddleware.RequestAuthHandker)

	v1.GET("/auth", func(c echo.Context) error { return c.NoContent(http.StatusNoContent) })

	order.RegisterRoutes(v1, ou)
	orderItem.RegisterRoutes(v1, oiu)
	adminUser.RegisterRoutes(v1, au)
	storeStaff.RegisterRoutes(v1, su)
	item.RegisterRoutes(v1, iu)
	report.RegisterRoutes(v1, ru)

	// routeの一覧表示
	for _, rr := range e.Routes() {
		fmt.Printf("%s %s\n", rr.Method, rr.Path)
	}

	return e
}
