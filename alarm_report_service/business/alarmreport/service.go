package alarmreport

import (
	"fmt"
)

type ReportAlarmRepo interface {
	FindReportAlarm(imei string, alarmTypeID []int, timeStart string, timeEnd string) (data []Domain, err error)
	CountData(imei string, timeStart string, timeEnd string, alarmType []int) int
}

type ReportAlarmService interface {
	FindReportAlarm(imei string, alarmTypeID []int, timeStart string, timeEnd string) (res Result, err error)
}
type reportAlarmService struct {
	reportAlarmRepo ReportAlarmRepo
}

func NewReportAlarmService(
	reportAlarmRepo ReportAlarmRepo,
) ReportAlarmService {
	return &reportAlarmService{
		reportAlarmRepo: reportAlarmRepo,
	}
}

func (c *reportAlarmService) FindReportAlarm(imei string, alarmTypeID []int, timeStart string, timeEnd string) (res Result, err error) {

	// offset := perPage*page - perPage

	r, err := c.reportAlarmRepo.FindReportAlarm(imei, alarmTypeID, timeStart, timeEnd)
	if err != nil {
		fmt.Println(err)
		return res, nil
	}
	// res.Domain = r
	// res.TotalAllData = c.reportAlarmRepo.CountData(imei, timeStart, timeEnd, alarmTypeID)
	// res.CurrentPage = page
	// res.PerPage = perPage
	// x := float64(res.TotalAllData) / float64(perPage)
	// res.TotalPage = int(math.Ceil(x))
	// return res, nil
	res.Domain = r
	res.TotalAllData = len(r)

	return res, nil

}
