package user

import (
	"authentication_service/business/user"
	"authentication_service/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) user.UserRepository {
	var userRepository user.UserRepository

	if dbCon.Driver == util.PostgreSQL {
		userRepository = NewPostgresRepository(dbCon.PostgreSQL)

	} else {
		panic("Database driver not supported")
	}

	return userRepository
}
