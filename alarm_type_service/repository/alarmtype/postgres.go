package alarmtype

import (
	"alarm_type_service/business/alarmtype"

	"gorm.io/gorm"
)

type AlarmTypeRepo struct {
	db *gorm.DB
}

func NewAlarmTypeRepository(db *gorm.DB) alarmtype.AlarmtypeRepo {
	return &AlarmTypeRepo{
		db: db,
	}
}

func (r *AlarmTypeRepo) ListAlarmtype(search string) (data []alarmtype.Domain) {
	record := []AlrmType{}
	res := r.db.Table("telegram.master_alert").Find(&record)
	if res.Error != nil {
		return []alarmtype.Domain{}
	}
	return toServiceList(record)
}
