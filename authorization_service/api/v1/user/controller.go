package user

import (
	"authorization_service/api/common/obj"
	_response "authorization_service/api/common/response"
	"authorization_service/api/v1/user/resp"
	service "authorization_service/business/user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(
	userService service.UserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (controller *UserController) GetPrivilege(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	user, err := controller.userService.FindUserPrivilege(intID)
	if err != nil {
		response := _response.BuildErrorResponse("User not found", obj.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}
	data := resp.FromService(*user)
	response := _response.BuildSuccsessResponse("User found", true, data)
	return c.JSON(http.StatusOK, response)
}
