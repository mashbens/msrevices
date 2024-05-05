package modules

import (
	"email_verifi_service/api"
	emailverifi "email_verifi_service/api/v1/emailverifi"
	emailverifiService "email_verifi_service/business/emailverifi"
	jwtService "email_verifi_service/business/jwt"
	"email_verifi_service/config"

	emailverifiRepo "email_verifi_service/repository/emailverifi"
	"email_verifi_service/util"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	emailverifiRepo := emailverifiRepo.RepositoryFactory(dbCon)
	jwtService := jwtService.NewJwtService()

	emailverifiService := emailverifiService.NewEmailVerifiService(emailverifiRepo)
	controller := api.Controller{
		Con: emailverifi.NewController(emailverifiService, jwtService),
	}
	return controller
}
