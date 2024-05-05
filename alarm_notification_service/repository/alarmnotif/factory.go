package alarmnotif

import (
	"alarm_notification_service/business/alarmnotif"
	"alarm_notification_service/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) alarmnotif.AlarmnotifRepo {
	var VehicleDetailtRepository alarmnotif.AlarmnotifRepo

	if dbCon.Driver == util.PostgreSQL {
		VehicleDetailtRepository = NewalarmnotifRepo(dbCon.PostgreSQL)

	} else {
		panic("Database driver not supported")
	}

	return VehicleDetailtRepository
}
