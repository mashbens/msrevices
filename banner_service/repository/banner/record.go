package banner

import (
	"banner_service/business/banner"
)

type Home_banner struct {
	Id          int
	Name        string
	File        string
	Status      int
	Position    int
	File_mobile string
	Count       int64
}

func (b *Home_banner) toService() banner.Domain {
	return banner.Domain{
		Id:          int(b.Id),
		Name:        b.Name,
		File:        b.File,
		Status:      b.Status,
		Position:    b.Position,
		File_mobile: b.File_mobile,
	}

}
func fromService(b banner.Domain) Home_banner {
	return Home_banner{
		Id:          int(b.Id),
		Name:        b.Name,
		File:        b.File,
		Status:      b.Status,
		Position:    b.Position,
		File_mobile: b.File_mobile,
	}
}

func toServiceList(data []Home_banner) []banner.Domain {
	a := []banner.Domain{}
	for key := range data {
		a = append(a, data[key].toService())
	}
	return a
}
