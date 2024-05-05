package fcmtoken

import (
	"time"
)

type FcmToken struct {
	ID         int
	Username   string
	Device_id  string
	Token      string
	Created_at time.Time
	Updated_at time.Time
}
