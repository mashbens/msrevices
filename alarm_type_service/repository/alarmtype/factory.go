package alarmtype

import (
	"alarm_type_service/business/alarmtype"
	"alarm_type_service/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) alarmtype.AlarmtypeRepo {
	var reportRepository alarmtype.AlarmtypeRepo

	if dbCon.Driver == util.PostgreSQL {
		reportRepository = NewAlarmTypeRepository(dbCon.PostgreSQL)

	} else {
		panic("Database driver not supported")
	}

	return reportRepository
}
