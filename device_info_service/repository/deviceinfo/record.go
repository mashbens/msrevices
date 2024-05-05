package deviceinfo

import (
	"device_info/business/deviceinfo"
	"time"
)

type Data struct {
	Imei        string
	Gsm_no      string
	Gps_type    int
	Veh_type    int
	Plate       string
	Device_name string
	Icon        int
	Year        int
	Stnk        time.Time
	Norangka    string
	Branch_Id   string
	SpeedLimit  int
	Nomesin     string
}

type GpsType struct {
	Name string
}

func (v *Data) toService() deviceinfo.Domain {
	return deviceinfo.Domain{
		Imei:        v.Imei,
		GsmNumber:   v.Gsm_no,
		GpsType:     v.Gps_type,
		VehicleType: v.Veh_type,
		Palte:       v.Plate,
		DeviceName:  v.Device_name,
		Icon:        v.Icon,
		Year:        v.Year,
		Stnk:        v.Stnk.Format("2006-01-02"),
		Norangka:    v.Norangka,
		Nomesin:     v.Nomesin,
		AdminBranch: v.Branch_Id,
		SpeedLimit:  v.SpeedLimit,
	}
}

//
