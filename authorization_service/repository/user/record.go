package user

import (
	"authorization_service/business/user"
)

type Users struct {
	ID        int
	Privilege int
}

func (u *Users) toService() user.Domain {
	return user.Domain{
		ID:        int64(u.ID),
		Privilege: u.Privilege,
	}
}

func fromService(u user.Domain) Users {
	return Users{
		ID:        int(u.ID),
		Privilege: u.Privilege,
	}
}
