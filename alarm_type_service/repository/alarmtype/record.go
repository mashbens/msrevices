package alarmtype

import (
	"alarm_type_service/business/alarmtype"
)

type AlrmType struct {
	ID       int
	Code     string
	Alert_no int
	Text_id  string
	Text_en  string
}

func (v *AlrmType) toService() alarmtype.Domain {
	return alarmtype.Domain{
		ID:       v.ID,
		Code:     v.Code,
		Alert_no: v.Alert_no,
		Text_id:  v.Text_id,
		Text_en:  v.Text_en,
	}
}

func toServiceList(data []AlrmType) []alarmtype.Domain {
	a := []alarmtype.Domain{}
	for key := range data {
		a = append(a, data[key].toService())
	}
	return a
}
