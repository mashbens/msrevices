package deviceinfo

import "fmt"

type DeviceInfoRepo interface {
	FindDeviceInfo(imei string) (data Domain, err error)
	FindGpsType(GpsNo int) string
	FindVehType(GpsNo int) string
}

type DeviceInfoService interface {
	FindDeviceInfo(imei string) (res Result, err error)
}

type deviceInfoService struct {
	deviceInfoRepo DeviceInfoRepo
}

func NewReportRunningService(
	deviceInfoRepo DeviceInfoRepo,
) DeviceInfoService {
	return &deviceInfoService{
		deviceInfoRepo: deviceInfoRepo,
	}
}

func (c *deviceInfoService) FindDeviceInfo(imei string) (res Result, err error) {
	data, err := c.deviceInfoRepo.FindDeviceInfo(imei)
	if err != nil {
		fmt.Println(err)
		return res, nil
	}
	// fmt.Println(data.GpsType)

	gpsType := c.deviceInfoRepo.FindGpsType(data.GpsType)
	vehType := c.deviceInfoRepo.FindVehType(data.VehicleType)
	data.StrVehicleType = vehType
	data.StrGpsType = gpsType
	res.Domain = data
	return res, nil
}
