package resp

import (
	"alarm_report_service/business/alarmreport"
)

type ReportResp struct {
	TotalAllData int
	// CurrentPage  int
	// PerPage      int
	// TotalPage    int
	Result []RespDomain `json:"result"`
}

type RespDomain struct {
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

func FromService(b alarmreport.Result) ReportResp {
	return ReportResp{
		TotalAllData: b.TotalAllData,
		// CurrentPage:  b.CurrentPage,
		// PerPage:      b.PerPage,
		// TotalPage:    b.TotalPage,
		Result: VehicleFromServiceSlice(b.Domain),
	}
}
func FromServiceDomain(v alarmreport.Domain) RespDomain {
	return RespDomain{
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

func VehicleFromServiceSlice(data []alarmreport.Domain) []RespDomain {
	var stopReportArray []RespDomain
	for key := range data {
		stopReportArray = append(stopReportArray, FromServiceDomain(data[key]))
	}
	return stopReportArray
}
