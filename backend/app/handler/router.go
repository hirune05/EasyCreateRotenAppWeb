package handler

import (
	"net/http"
	"time"
        "fmt"
        "os"

	adminUser "backend/app/handler/admin_user"
	"backend/app/handler/order"
	orderItem "backend/app/handler/order_item"
	storeStaff "backend/app/handler/store_staff"
        student "backend/app/handler/student"
        event "backend/app/handler/event"
        item "backend/app/handler/item"
	"backend/app/usecase"

	"backend/app/appmiddleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(ou usecase.Order, oiu usecase.OrderItem, au usecase.AdminUser, su usecase.StoreStaff, stu usecase.Student, eu usecase.Event, iu usecase.Item) http.Handler {
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

        student.RegisterRoutes(e.Group("/v1"), stu)
        event.RegisterRoutes(e.Group("/v1"), eu)
	v1 := e.Group("/v1")
        v1.Use(appmiddleware.RequestAuthHandker)

	order.RegisterRoutes(v1, ou)
	orderItem.RegisterRoutes(v1, oiu)
	adminUser.RegisterRoutes(v1, au)
	storeStaff.RegisterRoutes(v1, su)
        item.RegisterRoutes(v1, iu)

        // routeの一覧表示
        for _, rr := range e.Routes() {
          fmt.Printf("%s %s\n", rr.Method, rr.Path)
        }

	return e
}
