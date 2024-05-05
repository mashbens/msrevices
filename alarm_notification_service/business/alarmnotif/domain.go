package alarmnotif

import "time"

type Result struct {
	TotalAllData int
	CurrentPage  int
	PerPage      int
	TotalPage    int
	Domain       []Domain
}
type Domain struct {
	Id      int
	User_id int
	Type    string
	// Is_active bool
	Imei       string
	Username   string
	Alert_code string
	Body_id    string
	Body_en    string
	Time       time.Time
	StrTime    string
	Lon        float64
	Lat        float64
	StrLon     string
	StrLat     string
	Code       string
	Plat       string
}
