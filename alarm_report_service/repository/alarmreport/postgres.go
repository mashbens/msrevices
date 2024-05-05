package alarmreport

import (
	"alarm_report_service/business/alarmreport"
	"fmt"

	"gorm.io/gorm"
)

type ReportAlarmRepo struct {
	db *gorm.DB
}

func NewReportAlarmRepo(db *gorm.DB) alarmreport.ReportAlarmRepo {
	return &ReportAlarmRepo{
		db: db,
	}
}

func (r *ReportAlarmRepo) FindReportAlarm(imei string, alarmType []int, timeStart string, timeEnd string) (data []alarmreport.Domain, err error) {
	record := []ReportAlarm{}
	res := r.db.Order("time asc").Table("data.alert").Where("imei = ?", imei).Where("alert IN ?", alarmType).Where("time between ? and ?", timeStart, timeEnd).Find(&record)
	if res.Error != nil {
		fmt.Println(res.Error)
		return []alarmreport.Domain{}, res.Error
	}
	return toServiceList(record), nil

}
func (c *ReportAlarmRepo) CountData(imei string, timeStart string, timeEnd string, alarmType []int) int {
	var count int64
	res := c.db.Table("data.alert").Where("imei = ?", imei).Where("alert IN ?", alarmType).Where("time between ? and ?", timeStart, timeEnd).Order("time asc").Count(&count)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	return int(count)
}
