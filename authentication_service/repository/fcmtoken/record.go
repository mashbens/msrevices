package fcmtoken

import (
	"authentication_service/business/fcmtoken"
	"time"
)

type Fcm_token struct {
	ID         int
	Username   string
	Device_id  string
	Token      string
	Start_date time.Time
	Created_at time.Time
	Updated_at time.Time
}

func (b *Fcm_token) toServiceFcm() fcmtoken.FcmToken {
	return fcmtoken.FcmToken{
		ID:         int(b.ID),
		Username:   b.Username,
		Device_id:  b.Device_id,
		Token:      b.Token,
		Created_at: b.Created_at,
		Updated_at: b.Updated_at,
	}

}
