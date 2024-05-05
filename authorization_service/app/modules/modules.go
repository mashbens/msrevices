package modules

import (
	"authorization_service/api"
	"authorization_service/api/v1/user"
	userService "authorization_service/business/user"
	"authorization_service/config"
	userRepo "authorization_service/repository/user"
	"authorization_service/util"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	userRepo := userRepo.RepositoryFactory(dbCon)
	userService := userService.NewUserService(userRepo)
	controller := api.Controller{
		User: user.NewUserController(userService),
	}
	return controller
}
