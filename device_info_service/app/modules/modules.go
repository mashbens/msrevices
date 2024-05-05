package modules

import (
	"device_info/api"
	deviceInfo "device_info/api/v1/deviceinfo"
	deviceInfoService "device_info/business/deviceinfo"
	jwtService "device_info/business/jwt"
	"device_info/config"

	deviceInfoRepo "device_info/repository/deviceinfo"
	"device_info/util"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	deviceInfoRepo := deviceInfoRepo.RepositoryFactory(dbCon)
	jwtService := jwtService.NewJwtService()

	deviceInfoService := deviceInfoService.NewReportRunningService(deviceInfoRepo)
	controller := api.Controller{
		ReportRunning: deviceInfo.NewReportController(deviceInfoService, jwtService),
	}
	return controller
}
