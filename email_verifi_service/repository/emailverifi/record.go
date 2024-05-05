package emailverifi

import (
	"email_verifi_service/business/emailverifi"
	"time"
)

type Emailverifi struct {
	ID                 int
	Branch_id          int
	Privilage          int
	Fullname           string
	Company_name       string
	Email              string
	Phone              string
	Address            string
	Username           string
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

func (u *Emailverifi) toService() emailverifi.Domain {
	return emailverifi.Domain{
		ID:                 int(u.ID),
		Branch_id:          u.Branch_id,
		Privilage:          u.Privilage,
		Fullname:           u.Fullname,
		Company_name:       u.Company_name,
		Email:              u.Email,
		Phone:              u.Phone,
		Address:            u.Address,
		Username:           u.Username,
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

func toServiceDetails(data []Emailverifi) []emailverifi.Domain {
	a := []emailverifi.Domain{}
	for key := range data {
		a = append(a, data[key].toService())
	}
	return a
}
