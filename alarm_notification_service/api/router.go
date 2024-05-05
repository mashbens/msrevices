package api

import (
	"alarm_notification_service/api/v1/alarmnotif"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	Con *alarmnotif.Controller
}

func RegisterRoutes(e *echo.Echo, controller *Controller) {
	e.GET("", controller.Con.Alarmnotif)
}
