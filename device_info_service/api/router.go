package api

import (
	runningreport "device_info/api/v1/deviceinfo"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	ReportRunning *runningreport.ReportController
}

func RegisterRoutes(e *echo.Echo, controller *Controller) {
	e.GET("", controller.ReportRunning.Deviceinfo)
}
