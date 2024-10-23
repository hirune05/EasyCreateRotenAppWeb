package adminUser

import (
	"backend/app/usecase"

	"github.com/labstack/echo/v4"
)

type adminUserHandler struct {
	adminUserUseCase usecase.AdminUser
}

func RegisterRoutes(e *echo.Group, u usecase.AdminUser) {
	h := &adminUserHandler{
		adminUserUseCase: u,
	}

	e.POST("/adminUser", h.Create)
}
