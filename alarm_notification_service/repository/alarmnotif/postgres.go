package alarmnotif

import (
	"alarm_notification_service/business/alarmnotif"

	"gorm.io/gorm"
)

type alarmnotifRepo struct {
	db *gorm.DB
}

func NewalarmnotifRepo(db *gorm.DB) alarmnotif.AlarmnotifRepo {
	return &alarmnotifRepo{
		db: db,
	}
}

func (r *alarmnotifRepo) Findalarmnotif(username string, alertCode []string) (data []alarmnotif.Domain, err error) {
	record := []Alarmnotif{}
	res := r.db.Table("telegram.push_notification").Order("id desc").Where("username = ?", username).Where("alert_code IN ? ", alertCode).Find(&record)
	if res.Error != nil {
		return nil, res.Error
	}
	return toServiceList(record), nil
}

func (c *alarmnotifRepo) CountData(username string, alertCode []string) int {
	var count int64
	res := c.db.Table("telegram.push_notification").Where("username = ?", username).Where("alert_code IN ? ", alertCode).Count(&count)
	if res.Error != nil {
		return 0
	}
	return int(count)
}
func (c *alarmnotifRepo) FindPlate(imei string) string {
	var record Imei
	res := c.db.Table("data.gps").Where("imei = ?", imei).Find(&record)
	if res.Error != nil {
		return ""
	}
	return record.Plate
}
