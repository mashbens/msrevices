package fcmtoken

import (
	"authentication_service/business/fcmtoken"
	"authentication_service/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection2) fcmtoken.FcmTokenRepo {
	var fcmRepository fcmtoken.FcmTokenRepo

	if dbCon.Driver == util.PostgreSQL2 {
		fcmRepository = NewFcmTokenRepo(dbCon.PostgreSQL2)
	} else {
		panic("Database driver not supported")
	}

	return fcmRepository
}
