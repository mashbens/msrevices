package deviceinfo

import (
	"device_info/business/deviceinfo"

	"gorm.io/gorm"
)

type ReportDeviceInfoRepo struct {
	db *gorm.DB
}

func NewDeviceInfoRepo(db *gorm.DB) deviceinfo.DeviceInfoRepo {
	return &ReportDeviceInfoRepo{
		db: db,
	}
}

func (r *ReportDeviceInfoRepo) FindDeviceInfo(imei string) (data deviceinfo.Domain, err error) {
	var record Data
	res := r.db.Table("data.gps").Where("imei = ?", imei).Find(&record)
	if res.Error != nil {
		return data, res.Error
	}
	return record.toService(), nil
}
func (r *ReportDeviceInfoRepo) FindGpsType(id int) string {
	var record GpsType
	res := r.db.Table("master.gps_type").Where("id = ?", id).Find(&record)
	if res.Error != nil {
		return "data"
	}
	return record.Name
}
func (r *ReportDeviceInfoRepo) FindVehType(id int) string {
	var record GpsType
	res := r.db.Table("master.vehicle_type").Where("id = ?", id).Find(&record)
	if res.Error != nil {
		return "data"
	}
	return record.Name
}
