package alarmreport

import (
	"alarm_report_service/business/alarmreport"
	"alarm_report_service/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) alarmreport.ReportAlarmRepo {
	var ReportAlarmRepository alarmreport.ReportAlarmRepo

	if dbCon.Driver == util.PostgreSQL {
		ReportAlarmRepository = NewReportAlarmRepo(dbCon.PostgreSQL)

	} else {
		panic("Database driver not supported")
	}

	return ReportAlarmRepository
}
