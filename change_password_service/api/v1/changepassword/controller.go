package changepassword

import (
	"change_password_service/api/common/obj"
	_response "change_password_service/api/common/response"
	"change_password_service/api/v1/changepassword/request"
	"strings"

	service "change_password_service/business/changepassword"
	jwtService "change_password_service/business/jwt"
	"fmt"
	"strconv"

	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	Service    service.ChangepasswordService
	jwtService jwtService.JwtService
}

func NewController(
	Service service.ChangepasswordService,
	jwtService jwtService.JwtService,

) *Controller {
	return &Controller{
		Service:    Service,
		jwtService: jwtService,
	}
}
func (controller *Controller) Changepassword(c echo.Context) error {
	header := c.Request().Header.Get("Authorization")
	validate := controller.jwtService.ValidateToken(header)
	if header == "" {
		response := _response.BuildErrorResponse("Failed to validate token", obj.EmptyObj{})
		return c.JSON(http.StatusUnauthorized, response)
	}
	if validate != true {
		response := _response.BuildErrorResponse("Token expired or invalid", obj.EmptyObj{})
		return c.JSON(http.StatusUnauthorized, response)
	}
	tokenString := strings.Replace(header, "Bearer ", "", -1)
	token, _ := jwt.Parse(tokenString, nil)
	if token == nil {
		return nil
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	intID, _ := strconv.Atoi(id)

	payload := new(request.Req)
	err := c.Bind(payload)
	if err != nil {
		response := _response.BuildErrorResponse("Invalid request body", obj.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	res, err := controller.Service.ChangePassword(intID, payload.OldPass, payload.NewPass)
	fmt.Println(err, "err handler")
	if err != nil {
		response := _response.BuildErrorResponse("There was a problem with your update.", obj.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}
	_ = res
	// data := resp.FromService(res)

	_response := _response.BuildSuccsessResponse("Succsess change password", true, obj.EmptyObj{})
	return c.JSON(http.StatusOK, _response)
}
