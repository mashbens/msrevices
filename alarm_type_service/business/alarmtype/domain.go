package alarmtype

type Result struct {
	TotalAllData int
	CurrentPage  int
	PerPage      int
	TotalPage    int
	Domain       []Domain
}

type Domain struct {
	ID         int
	Code       string
	Alert_no   int
	Text_id    string
	Text_en    string
	Checked    bool
	Alert_code string
}
