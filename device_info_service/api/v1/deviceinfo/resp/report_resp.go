package resp

import (
	"device_info/business/deviceinfo"
)

type ReportResp struct {
	Result RespDomain `json:"result"`
}

type RespDomain struct {
	Imei            string
	GpsType         string
	DeviceName      string
	VehicleType     string
	Palte           string
	GsmNumber       string
	Icon            int
	Stnk            string
	AdminBranch     string
	OwnerName       string
	InstalationTech string
	Year            int
	MachineNumber   string
	ChassisNumber   string
	SpeedLimit      int
	Norangka        string
	Nomesin         string
}

func FromService(b deviceinfo.Result) ReportResp {
	return ReportResp{
		Result: FromServiceDomain(b.Domain),
	}
}
func FromServiceDomain(v deviceinfo.Domain) RespDomain {
	return RespDomain{
		Imei:        v.Imei,
		GsmNumber:   v.GsmNumber,
		GpsType:     v.StrGpsType,
		VehicleType: v.StrVehicleType,
		Palte:       v.Palte,
		DeviceName:  v.DeviceName,
		Icon:        v.Icon,
		Year:        v.Year,
		Stnk:        v.Stnk,
		Norangka:    v.Norangka,
		Nomesin:     v.Nomesin,
		AdminBranch: v.AdminBranch,
		SpeedLimit:  v.SpeedLimit,
	}
}

// func VehicleFromServiceSlice(data []deviceinfo.Domain) []RespDomain {
// 	var stopReportArray []RespDomain
// 	for key := range data {
// 		stopReportArray = append(stopReportArray, FromServiceDomain(data[key]))
// 	}
// 	return stopReportArray
// }
