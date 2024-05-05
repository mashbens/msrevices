package api

import (
	"authorization_service/api/v1/user"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	User *user.UserController
}

func RegisterRoutes(e *echo.Echo, controller *Controller) {
	userRoutes := e.Group("/AuthenticationUser/")
	userRoutes.GET(":id", controller.User.GetPrivilege)

}
