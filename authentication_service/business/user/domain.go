package user

import (
	"time"
)

type Domain struct {
	ID                 int64
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
	Is_verifed_email   bool
	Is_verifed_pgone   bool
	Type_customer      int
	Prename            string
	Apikey             string
	Token              string
	FcmToken           string
	OldFcmToken        string
	NewFcmToken        string
}

type FcmToken struct {
	ID         int
	Username   string
	Privilage  int
	Device_id  string
	Token      string
	Created_at time.Time
	Updated_at time.Time
}
