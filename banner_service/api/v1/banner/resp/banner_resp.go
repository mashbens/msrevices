package resp

import (
	"banner_service/business/banner"
)

type BannerResp struct {
	TotalAllData int
	CurrentPage  int
	PerPage      int
	TotalPage    int
	Result       []BannerRespDomain
}
type BannerRespDomain struct {
	Id          int
	Name        string
	File        string
	Status      int
	Position    int
	File_mobile string
}

func FromService(b banner.Result) BannerResp {
	return BannerResp{
		TotalAllData: b.TotalAllData,
		CurrentPage:  b.CurrentPage,
		PerPage:      b.PerPage,
		TotalPage:    b.TotalPage,
		Result:       FromServiceSlice(b.Domain),
	}
}

func FromServiceDomain(b banner.Domain) BannerRespDomain {
	return BannerRespDomain{
		Id:          b.Id,
		Name:        b.Name,
		File:        b.File,
		Status:      b.Status,
		Position:    b.Position,
		File_mobile: b.File_mobile,
	}
}
func FromServiceSlice(data []banner.Domain) []BannerRespDomain {
	var bannerArray []BannerRespDomain
	for key := range data {
		bannerArray = append(bannerArray, FromServiceDomain(data[key]))

	}
	return bannerArray
}
