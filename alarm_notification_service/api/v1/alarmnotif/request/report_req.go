package request

type AlarmtRequest struct {
	Username    string   `json:"username"`
	AlarmTypeID []string `json:"alrmTypeID"`
}
