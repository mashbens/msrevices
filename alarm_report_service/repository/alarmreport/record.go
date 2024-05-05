package alarmreport

import (
	"alarm_report_service/business/alarmreport"
)

type ReportAlarm struct {
	ID           int
	Imei         string
	Lon          float64
	Lat          float64
	Speed        int
	Address      string
	Time         string
	Alert        int
	Remark_notif int
	Remark_email int
	Alert_text   string
}

func (v *ReportAlarm) toService() alarmreport.Domain {
	return alarmreport.Domain{
		Imei:         v.Imei,
		ID:           v.ID,
		Lon:          v.Lon,
		Lat:          v.Lat,
		Speed:        v.Speed,
		Address:      v.Address,
		Time:         v.Time,
		Alert:        v.Alert,
		Alert_text:   v.Alert_text,
		Remark_notif: v.Remark_notif,
		Remark_email: v.Remark_email,
	}
}

func toServiceList(data []ReportAlarm) []alarmreport.Domain {
	a := []alarmreport.Domain{}
	for key := range data {
		a = append(a, data[key].toService())
	}
	return a
}
