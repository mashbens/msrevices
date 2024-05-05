package alarmnotif

import (
	"fmt"
	"strings"
)

type AlarmnotifRepo interface {
	Findalarmnotif(username string, alertCode []string) (data []Domain, err error)
	CountData(username string, alertCode []string) int
	FindPlate(imei string) string
}

type AlarmnotifService interface {
	Findalarmnotif(userID int, username string, alertCode []string) (res Result, err error)
}

type alarmnotifService struct {
	alarmnotifRepo AlarmnotifRepo
}

func NewalarmnotifService(
	alarmnotifRepo AlarmnotifRepo,
) AlarmnotifService {
	return &alarmnotifService{
		alarmnotifRepo: alarmnotifRepo,
	}
}

func (c *alarmnotifService) Findalarmnotif(userID int, username string, alertCode []string) (res Result, err error) {
	list, err := c.alarmnotifRepo.Findalarmnotif(username, alertCode)
	if err != nil {
		fmt.Println(err)
		return res, err
	}
	domain := []Domain{}
	for _, r := range list {
		r.Plat = c.alarmnotifRepo.FindPlate(r.Imei)
		if r.Lon == 0 {
			r.StrLon = "0.0"
		} else {
			sLon := fmt.Sprintf("%f", r.Lon)
			r.StrLon = sLon
		}
		if r.Lat == 0 {
			r.StrLat = "0.0"
		} else {
			sLat := fmt.Sprintf("%f", r.Lat)
			r.StrLat = sLat
		}
		r.StrTime = r.Time.Format("2006-01-02 15:04")
		r.Code = r.Alert_code
		s := strings.ReplaceAll(r.Code, "_", " ")
		r.Code = strings.Title(strings.ToLower(s))
		domain = append(domain, r)
	}

	res.Domain = domain

	res.TotalAllData = len(domain)

	return res, nil
}
