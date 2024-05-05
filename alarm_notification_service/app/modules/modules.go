package modules

import (
	"alarm_notification_service/api"
	alarmnotif "alarm_notification_service/api/v1/alarmnotif"
	alarmnotifService "alarm_notification_service/business/alarmnotif"
	jwtService "alarm_notification_service/business/jwt"
	"alarm_notification_service/config"

	alarmnotifRepo "alarm_notification_service/repository/alarmnotif"
	"alarm_notification_service/util"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	alarmnotifRepo := alarmnotifRepo.RepositoryFactory(dbCon)
	jwtService := jwtService.NewJwtService()

	alarmnotifService := alarmnotifService.NewalarmnotifService(alarmnotifRepo)
	controller := api.Controller{
		Con: alarmnotif.NewController(alarmnotifService, jwtService),
	}
	return controller
}
