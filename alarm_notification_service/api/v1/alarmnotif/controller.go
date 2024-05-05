package alarmnotif

import (
	"alarm_notification_service/api/common/obj"
	_response "alarm_notification_service/api/common/response"
	"alarm_notification_service/api/v1/alarmnotif/request"
	"alarm_notification_service/api/v1/alarmnotif/resp"
	service "alarm_notification_service/business/alarmnotif"
	jwtService "alarm_notification_service/business/jwt"

	"fmt"
	"strconv"
	"strings"

	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	Service    service.AlarmnotifService
	jwtService jwtService.JwtService
}

func NewController(
	Service service.AlarmnotifService,
	jwtService jwtService.JwtService,

) *Controller {
	return &Controller{
		Service:    Service,
		jwtService: jwtService,
	}
}
func (controller *Controller) Alarmnotif(c echo.Context) error {
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

	payload := new(request.AlarmtRequest)
	err := c.Bind(payload)

	claims, _ := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	intID, _ := strconv.Atoi(id)
	res, err := controller.Service.Findalarmnotif(intID, payload.Username, payload.AlarmTypeID)
	if err != nil {
		response := _response.BuildErrorResponse("Data not found.", obj.EmptyObj{})
		return c.JSON(http.StatusInternalServerError, response)
	}
	data := resp.FromService(res)

	_response := _response.BuildSuccsessResponse("Succsess get data", true, data)
	return c.JSON(http.StatusOK, _response)
}
