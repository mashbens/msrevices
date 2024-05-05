package alarmtype

import (
	"strings"
)

type AlarmtypeRepo interface {
	ListAlarmtype(search string) (data []Domain)
}

type AlarmtypeService interface {
	ListAlarmtype(search string) (res Result, err error)
}

type alarmtypeService struct {
	alarmtypeRepo AlarmtypeRepo
}

func NewAlarmtypeService(
	alarmtypeRepo AlarmtypeRepo,
) AlarmtypeService {
	return &alarmtypeService{
		alarmtypeRepo: alarmtypeRepo,
	}
}

func (c *alarmtypeService) ListAlarmtype(search string) (res Result, err error) {
	data := c.alarmtypeRepo.ListAlarmtype(search)
	domain := []Domain{}
	for _, d := range data {
		d.Alert_code = d.Code
		s := strings.ReplaceAll(d.Code, "_", " ")
		d.Code = strings.Title(strings.ToLower(s))
		domain = append(domain, d)
	}
	res.Domain = domain
	// fmt.Println(res)
	return res, nil
}
