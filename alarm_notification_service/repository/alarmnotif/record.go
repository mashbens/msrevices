package alarmnotif

import (
	"alarm_notification_service/business/alarmnotif"
	"time"
)

type Alarmnotif struct {
	Id      int
	User_id int
	Type    string
	// Is_active bool
	Imei       string
	Username   string
	Alert_code string
	Body_id    string
	Body_en    string
	Updated_at time.Time
	Lon        float64
	Lat        float64
	Plat       string
}

func (v *Alarmnotif) toService() alarmnotif.Domain {
	return alarmnotif.Domain{
		Id:      v.Id,
		User_id: v.User_id,
		Type:    v.Type,
		// Is_active: v.Is_active,
		Imei:       v.Imei,
		Username:   v.Username,
		Alert_code: v.Alert_code,
		Body_id:    v.Body_id,
		Body_en:    v.Body_en,
		Time:       v.Updated_at,
		// Time:       v.Time,
		Lon:  v.Lon,
		Lat:  v.Lat,
		Plat: v.Plat,
	}
}

func toServiceList(data []Alarmnotif) []alarmnotif.Domain {
	a := []alarmnotif.Domain{}
	for key := range data {
		a = append(a, data[key].toService())
	}
	return a
}

type Imei struct {
	Plate string
}
