package resp

import (
	"alarm_type_service/business/alarmtype"
)

type ReportResp struct {
	// TotalAllData int
	// CurrentPage  int
	// PerPage      int
	// TotalPage    int
	Result []RespDomain `json:"result"`
}

type RespDomain struct {
	// ID       int
	Code     string
	Alert_no int
	// Text_id  string
	// Text_en  string
	Alert_code string
	Checked    bool
}

func FromService(b alarmtype.Result) ReportResp {
	return ReportResp{
		// TotalAllData: b.TotalAllData,
		// CurrentPage:  b.CurrentPage,
		// PerPage:      b.PerPage,
		// TotalPage:    b.TotalPage,
		Result: VehicleFromServiceSlice(b.Domain),
	}
}
func FromServiceDomain(v alarmtype.Domain) RespDomain {
	return RespDomain{
		// ID:       v.ID,
		Code:     v.Code,
		Alert_no: v.Alert_no,
		// Text_id:  v.Text_id,
		// Text_en:  v.Text_en,
		Alert_code: v.Alert_code,
		Checked:    v.Checked,
	}
}

func VehicleFromServiceSlice(data []alarmtype.Domain) []RespDomain {
	var stopReportArray []RespDomain
	for key := range data {
		stopReportArray = append(stopReportArray, FromServiceDomain(data[key]))
	}
	return stopReportArray
}
