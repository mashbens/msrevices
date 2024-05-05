package deviceinfo

import (
	"device_info/business/deviceinfo"
	"device_info/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) deviceinfo.DeviceInfoRepo {
	var ReportRunningRepository deviceinfo.DeviceInfoRepo

	if dbCon.Driver == util.PostgreSQL {
		ReportRunningRepository = NewDeviceInfoRepo(dbCon.PostgreSQL)

	} else {
		panic("Database driver not supported")
	}

	return ReportRunningRepository
}
