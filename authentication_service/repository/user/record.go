package user

import (
	"authentication_service/business/user"
	"time"
)

type Users struct {
	ID                 int
	Branch_id          int
	Privilage          int
	Fullname           string
	Company_name       string
	Email              string
	Phone              string
	Address            string
	Username           string
	Password           string
	Status             int
	Last_login         time.Time
	Forgot_key         string
	Created_at         time.Time
	Created_by         int
	Updated_at         time.Time
	Updated_by         int
	Deleted_at         time.Time
	Deleted_by         int
	Type               int
	Remark             string
	Email_verification string
}

type Fcm_token struct {
	ID         int
	Username   string
	Device_id  string
	Token      string
	Start_date time.Time
	Created_at time.Time
	Updated_at time.Time
}

func (b *Fcm_token) toServiceFcm() user.FcmToken {
	return user.FcmToken{
		ID:         int(b.ID),
		Username:   b.Username,
		Device_id:  b.Device_id,
		Token:      b.Token,
		Created_at: b.Created_at,
		Updated_at: b.Updated_at,
	}

}
func (u *Users) toService() user.Domain {
	return user.Domain{
		ID:                 int64(u.ID),
		Branch_id:          u.Branch_id,
		Privilage:          u.Privilage,
		Fullname:           u.Fullname,
		Company_name:       u.Company_name,
		Email:              u.Email,
		Phone:              u.Phone,
		Address:            u.Address,
		Username:           u.Username,
		Password:           u.Password,
		Status:             u.Status,
		Last_login:         u.Last_login,
		Forgot_key:         u.Forgot_key,
		Created_at:         u.Created_at,
		Created_by:         u.Created_by,
		Updated_at:         u.Updated_at,
		Updated_by:         u.Updated_by,
		Deleted_at:         u.Deleted_at,
		Deleted_by:         u.Deleted_by,
		Type:               u.Type,
		Remark:             u.Remark,
		Email_verification: u.Email_verification,
	}
}

func fromService(u user.Domain) Users {
	return Users{
		ID:                 int(u.ID),
		Branch_id:          u.Branch_id,
		Privilage:          u.Privilage,
		Fullname:           u.Fullname,
		Company_name:       u.Company_name,
		Email:              u.Email,
		Phone:              u.Phone,
		Address:            u.Address,
		Username:           u.Username,
		Password:           u.Password,
		Status:             u.Status,
		Last_login:         u.Last_login,
		Forgot_key:         u.Forgot_key,
		Created_at:         u.Created_at,
		Created_by:         u.Created_by,
		Updated_at:         u.Updated_at,
		Updated_by:         u.Updated_by,
		Deleted_at:         u.Deleted_at,
		Deleted_by:         u.Deleted_by,
		Type:               u.Type,
		Remark:             u.Remark,
		Email_verification: u.Email_verification,
	}
}
