package alarmreport

type Result struct {
	TotalAllData int
	CurrentPage  int
	PerPage      int
	TotalPage    int
	Domain       []Domain
}

type Domain struct {
	ID           int
	Imei         string
	Lon          float64
	Lat          float64
	Speed        int
	Address      string
	Time         string
	Alert        int
	Remark_notif int
	Remark_email int
	Alert_text   string
}
