package api

import (
	"alarm_type_service/api/v1/alarmtype"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	Report *alarmtype.ReportController
}

func RegisterRoutes(e *echo.Echo, controller *Controller) {
	e.GET("", controller.Report.ReportParking)
}
