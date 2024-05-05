package emailverifi

import (
	"email_verifi_service/api/common/obj"
	_response "email_verifi_service/api/common/response"
	"email_verifi_service/api/v1/emailverifi/resp"
	service "email_verifi_service/business/emailverifi"
	jwtService "email_verifi_service/business/jwt"
	"fmt"
	"strconv"

	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	Service    service.EmailVerifiService
	jwtService jwtService.JwtService
}

func NewController(
	Service service.EmailVerifiService,
	jwtService jwtService.JwtService,

) *Controller {
	return &Controller{
		Service:    Service,
		jwtService: jwtService,
	}
}
func (controller *Controller) Emailverifi(c echo.Context) error {
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
	tokenString := header
	token, _ := jwt.Parse(tokenString, nil)
	if token == nil {
		return nil
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	intID, _ := strconv.Atoi(id)
	res, err := controller.Service.FindEmailVerifi(intID)
	if err != nil {
		response := _response.BuildErrorResponse("Data not found.", obj.EmptyObj{})
		return c.JSON(http.StatusInternalServerError, response)
	}
	data := resp.FromService(res)

	_response := _response.BuildSuccsessResponse("Succsess get data", true, data)
	return c.JSON(http.StatusOK, _response)
}
