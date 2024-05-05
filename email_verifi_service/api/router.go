package api

import (
	"email_verifi_service/api/v1/emailverifi"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	Con *emailverifi.Controller
}

func RegisterRoutes(e *echo.Echo, controller *Controller) {
	e.GET("", controller.Con.Emailverifi)
}
