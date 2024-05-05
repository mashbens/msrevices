package modules

import (
	"banner_service/api"
	"banner_service/api/v1/banner"
	bannerService "banner_service/business/banner"
	jwtService "banner_service/business/jwt"
	"banner_service/config"
	bannerRepo "banner_service/repository/banner"
	"banner_service/util"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	bannerRepo := bannerRepo.RepositoryFactory(dbCon)
	bannerService := bannerService.NewBannerService(bannerRepo)
	jwtService := jwtService.NewJwtService()
	controller := api.Controller{
		Banner: banner.NewBannerController(bannerService, jwtService),
	}
	return controller
}
