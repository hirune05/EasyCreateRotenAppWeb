package item

import (
	"backend/app/usecase"

	"github.com/labstack/echo/v4"
)

type reportHandler struct {
	reportUseCase usecase.Report
}

func RegisterRoutes(e *echo.Group, u usecase.Report) {
	h := &reportHandler{
		reportUseCase: u,
	}
	e.POST("/report-test", h.SendEmail)
}
