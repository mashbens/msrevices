package banner

import (
	"banner_service/business/banner"
	"banner_service/util"
)

func RepositoryFactory(dbCon *util.DatabaseConnection) banner.BannerRepo {
	var bannerRepository banner.BannerRepo

	if dbCon.Driver == util.PostgreSQL {
		bannerRepository = NewBannerRepository(dbCon.PostgreSQL)

	} else {
		panic("Database driver not supported")
	}

	return bannerRepository
}
