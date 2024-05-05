package api

import (
	reportstop "alarm_report_service/api/v1/alarmreport"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	ReportStop *reportstop.ReportController
}

func RegisterRoutes(e *echo.Echo, controller *Controller) {
	e.GET("", controller.ReportStop.ReportParking)
}
