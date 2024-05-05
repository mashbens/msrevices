package api

import (
	"change_password_service/api/v1/changepassword"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	Con *changepassword.Controller
}

func RegisterRoutes(e *echo.Echo, controller *Controller) {
	e.POST("", controller.Con.Changepassword)
}
