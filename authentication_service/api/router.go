package api

import (
	"authentication_service/api/v1/auth"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	Auth *auth.AuthController
}

func RegisterRoutes(e *echo.Echo, controller *Controller) {
	e.POST("login", controller.Auth.LoginHandller)
	e.POST("logout", controller.Auth.Logout)
}
