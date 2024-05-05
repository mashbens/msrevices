package modules

import (
	"alarm_type_service/api"
	alarmtype "alarm_type_service/api/v1/alarmtype"
	alarmtypeService "alarm_type_service/business/alarmtype"
	jwtService "alarm_type_service/business/jwt"
	"alarm_type_service/config"

	alarmtypeRepo "alarm_type_service/repository/alarmtype"
	"alarm_type_service/util"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	alarmtypeRepo := alarmtypeRepo.RepositoryFactory(dbCon)
	jwtService := jwtService.NewJwtService()

	alarmtypeService := alarmtypeService.NewAlarmtypeService(alarmtypeRepo)
	controller := api.Controller{
		Report: alarmtype.NewReportController(alarmtypeService, jwtService),
	}
	return controller
}
