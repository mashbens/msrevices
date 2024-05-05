package request

import "authentication_service/business/user"

type DeleteFcmReq struct {
	FcmToken string `json:"fcm_token"`
}

func NewDeleteFcmReq(req DeleteFcmReq) user.Domain {
	return user.Domain{
		FcmToken: req.FcmToken,
	}
}
