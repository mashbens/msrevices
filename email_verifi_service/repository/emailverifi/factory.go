package emailverifi

import (
	"email_verifi_service/business/emailverifi"
	"email_verifi_service/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) emailverifi.EmailVerifiRepo {
	var Repository emailverifi.EmailVerifiRepo

	if dbCon.Driver == util.PostgreSQL {
		Repository = NewemailverifiRepo(dbCon.PostgreSQL)

	} else {
		panic("Database driver not supported")
	}

	return Repository
}
