package resp

import (
	"alarm_notification_service/business/alarmnotif"
)

type Resp struct {
	TotalAllData int

	Result []RespDomain `json:"result"`
}

type RespDomain struct {
	User_id int
	Type    string
	// Is_active bool
	Imei       string
	Username   string
	Plat       string
	Code       string
	Alert_code string
	Body_id    string
	Body_en    string
	Time       string
	Lon        string
	Lat        string
}

func FromService(b alarmnotif.Result) Resp {
	return Resp{
		TotalAllData: b.TotalAllData,

		Result: FromServiceSlice(b.Domain),
	}
}
func FromServiceDomain(v alarmnotif.Domain) RespDomain {
	return RespDomain{
		User_id: v.User_id,
		Type:    v.Type,
		// Is_active: v.Is_active,
		Imei:       v.Imei,
		Username:   v.Username,
		Alert_code: v.Alert_code,
		Body_id:    v.Body_id,
		Body_en:    v.Body_en,
		Time:       v.StrTime,
		Lon:        v.StrLon,
		Lat:        v.StrLat,
		Code:       v.Code,
		Plat:       v.Plat,
	}
}
func FromServiceSlice(data []alarmnotif.Domain) []RespDomain {
	var stopReportArray []RespDomain
	for key := range data {
		stopReportArray = append(stopReportArray, FromServiceDomain(data[key]))
	}
	return stopReportArray
}
