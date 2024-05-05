package request

type ReportRequest struct {
	Imei      string `json:"imei"`
	TimeStart string `json:"timeStart"`
	TimeEnd   string `json:"timeEnd"`
}
