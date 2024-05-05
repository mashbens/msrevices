package deviceinfo

import (
	"device_info/api/common/obj"
	_response "device_info/api/common/response"
	"device_info/api/v1/deviceinfo/resp"
	service "device_info/business/deviceinfo"
	jwtService "device_info/business/jwt"
	"fmt"

	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/golang-jwt/jwt"
)

type ReportController struct {
	ReportService service.DeviceInfoService
	jwtService    jwtService.JwtService
}

func NewReportController(
	ReportService service.DeviceInfoService,
	jwtService jwtService.JwtService,

) *ReportController {
	return &ReportController{
		ReportService: ReportService,
		jwtService:    jwtService,
	}
}
func (controller *ReportController) Deviceinfo(c echo.Context) error {
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

	imei := c.Request().URL.Query().Get("imei")

	res, err := controller.ReportService.FindDeviceInfo(imei)
	if err != nil {
		fmt.Println(err)
		response := _response.BuildErrorResponse("Internal Server error", obj.EmptyObj{})
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := resp.FromService(res)
	_response := _response.BuildSuccsessResponse("Succsess get data", true, data)
	return c.JSON(http.StatusOK, _response)

}
