package alarmtype

import (
	"alarm_type_service/api/common/obj"
	_response "alarm_type_service/api/common/response"
	"alarm_type_service/api/v1/alarmtype/request"
	"alarm_type_service/api/v1/alarmtype/resp"
	service "alarm_type_service/business/alarmtype"
	jwtService "alarm_type_service/business/jwt"
	"fmt"
	"time"

	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/golang-jwt/jwt"
)

type ReportController struct {
	ReportService service.AlarmtypeService
	jwtService    jwtService.JwtService
}

func NewReportController(
	ReportService service.AlarmtypeService,
	jwtService jwtService.JwtService,

) *ReportController {
	return &ReportController{
		ReportService: ReportService,
		jwtService:    jwtService,
	}
}
func (controller *ReportController) ReportParking(c echo.Context) error {
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
	_ = intID

	payload := new(request.ReportRequest)
	err := c.Bind(payload)

	if err != nil {
		response := _response.BuildErrorResponse("Invalid request body", obj.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	layout := "2006-01-02 15:04:05 -0700 MST"

	vTimeStart := payload.TimeStart + ":03 +0000 UTC"
	vTimeEnd := payload.TimeEnd + ":03 +0000 UTC"

	t1, err := time.Parse(layout, vTimeStart)
	t2, err := time.Parse(layout, vTimeEnd)

	if t2.Unix()-t1.Unix() > 259200 {
		response := _response.BuildErrorResponse("The time interval must be 3 days", obj.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	res, err := controller.ReportService.ListAlarmtype("")
	if err != nil {
		response := _response.BuildErrorResponse("Internal Server error", obj.EmptyObj{})
		return c.JSON(http.StatusInternalServerError, response)
	}
	if cap(res.Domain) == 0 {
		_response := _response.BuildSuccsessResponse("Succsess get data", true, obj.EmptyObj{})
		return c.JSON(http.StatusOK, _response)
	}
	data := resp.FromService(res)
	_response := _response.BuildSuccsessResponse("Succsess get data", true, data)
	return c.JSON(http.StatusOK, _response)

}
