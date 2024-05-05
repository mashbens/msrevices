package auth

import (
	"authentication_service/api/common/obj"
	_response "authentication_service/api/common/response"
	"authentication_service/api/v1/auth/request"
	"authentication_service/api/v1/auth/resp"
	service "authentication_service/business/auth"
	fcmTokenService "authentication_service/business/fcmtoken"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authService     service.AuthService
	fcmTokenService fcmTokenService.FcmTokenService
}

func NewAuthController(
	authService service.AuthService,
	fcmTokenService fcmTokenService.FcmTokenService,

) *AuthController {
	return &AuthController{
		authService:     authService,
		fcmTokenService: fcmTokenService,
	}
}

func (controller *AuthController) LoginHandller(c echo.Context) error {
	var loginRequest request.LoginRequest
	err := c.Bind(&loginRequest)
	if err != nil {
		response := _response.BuildErrorResponse("Invalid request body", obj.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}

	if loginRequest.Username == "" || loginRequest.Password == "" {
		response := _response.BuildErrorResponse("Invalid request body", obj.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}
	user, err := controller.authService.Login(request.NewLoginRequest(loginRequest))
	if err != nil {
		response := _response.BuildErrorResponse("There was a problem with your login.", obj.EmptyObj{})
		return c.JSON(http.StatusUnauthorized, response)
	}
	if user.Token == "" {
		response := _response.BuildErrorResponse("There was a problem with your login.", obj.EmptyObj{})
		return c.JSON(http.StatusUnauthorized, response)
	}
	c.Response().Header().Set("Token", user.Token)
	data := resp.FromService(*user)
	_response := _response.BuildSuccsessResponse("User logged in successfully", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *AuthController) Logout(c echo.Context) error {
	header := c.Request().Header.Get("Authorization")
	if header == "" {
		response := _response.BuildErrorResponse("Failed to validate token", obj.EmptyObj{})
		log.Println("Failed to validate token")
		return c.JSON(http.StatusUnauthorized, response)
	}
	tokenString := strings.Replace(header, "Bearer ", "", -1)

	if tokenString == "" {
		response := _response.BuildErrorResponse("Invalid request body", obj.EmptyObj{})
		return c.JSON(http.StatusBadRequest, response)
	}
	fcm := controller.fcmTokenService.DeleteFcmToken(tokenString)
	if fcm != nil {
		fmt.Println(fcm)
	}

	// data := resp.FromService(*user)
	_response := _response.BuildSuccsessResponse("Succsess logout & Fcm token deleted", true, obj.EmptyObj{})
	return c.JSON(http.StatusOK, _response)
}
