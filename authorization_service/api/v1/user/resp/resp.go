package resp

import (
	"authorization_service/business/user"
)

type UserResp struct {
	ID        int `json:"ID"`
	Privilege int `json:"Privilege"`
}

func FromService(u user.Domain) UserResp {
	return UserResp{
		ID:        int(u.ID),
		Privilege: u.Privilege,
	}
}
