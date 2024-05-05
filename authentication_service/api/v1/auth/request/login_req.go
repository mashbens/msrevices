package request

import "authentication_service/business/user"

type LoginRequest struct {
	Username string `json:"username" form:"username" `
	Password string `json:"password" form:"password" binding:"required,min=6"`
	FcmToken string `json:"fcm_token" `
}

func NewLoginRequest(req LoginRequest) user.Domain {
	return user.Domain{
		Username: req.Username,
		Password: req.Password,
		FcmToken: req.FcmToken,
	}
}
