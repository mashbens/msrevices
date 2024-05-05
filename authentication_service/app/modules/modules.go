package modules

import (
	"authentication_service/api"
	"authentication_service/api/v1/auth"
	authService "authentication_service/business/auth"
	authorizationService "authentication_service/business/authorization"
	fcmTokenService "authentication_service/business/fcmtoken"
	jwtService "authentication_service/business/jwt"
	userService "authentication_service/business/user"
	"authentication_service/config"
	fcmTokenRepo "authentication_service/repository/fcmtoken"
	userRepo "authentication_service/repository/user"
	"authentication_service/util"
)

func RegisterModules(dbCon *util.DatabaseConnection, dbCon2 *util.DatabaseConnection2, config *config.AppConfig) api.Controller {
	userRepo := userRepo.RepositoryFactory(dbCon)
	jwtService := jwtService.NewAuthService()
	authorizationService := authorizationService.NewAuthService()
	userService := userService.NewUserService(userRepo)

	fcmTokenRepo := fcmTokenRepo.RepositoryFactory(dbCon2)

	fcmTokenService := fcmTokenService.NewFcmService(fcmTokenRepo)
	authService := authService.NewAuthService(userService, authorizationService, jwtService, fcmTokenService)
	controller := api.Controller{
		Auth: auth.NewAuthController(authService, fcmTokenService),
		// User: user.NewUserController(userService, jwtService),
	}
	return controller
}
