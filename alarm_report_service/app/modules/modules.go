package modules

import (
	"alarm_report_service/api"
	alarmreport "alarm_report_service/api/v1/alarmreport"
	alarmreportService "alarm_report_service/business/alarmreport"
	jwtService "alarm_report_service/business/jwt"
	"alarm_report_service/config"

	alarmreportRepo "alarm_report_service/repository/alarmreport"
	"alarm_report_service/util"
)

func RegisterModules(dbCon *util.DatabaseConnection, config *config.AppConfig) api.Controller {
	alarmreportRepo := alarmreportRepo.RepositoryFactory(dbCon)
	jwtService := jwtService.NewJwtService()

	alarmreportService := alarmreportService.NewReportAlarmService(alarmreportRepo)
	controller := api.Controller{
		ReportStop: alarmreport.NewReportController(alarmreportService, jwtService),
	}
	return controller
}
