package resp

import (
	"authentication_service/business/user"
)

type TokenClaims struct {
	UserID int
}
type UserResp struct {
	ID           int    `json:"ID"`
	Username     string `json:"Username"`
	Fullname     string `json:"Fullname"`
	Phone        string `json:"Phone"`
	Email        string `json:"Email"`
	Addres       string `json:"Addres"`
	Company_name string `json:"Company_name"`
	Branch_id    int    `json:"Branch_id"`
	Privilage    int    `json:"Privilage"`
	// Password     string `json:"Password"`
	// Status             int       `json:"Status"`
	// Last_login         time.Time `json:"Last_login"`
	// Forgot_key         string    `json:"Forgot_key"`
	// Created_at         time.Time `json:"Created_at"`
	// Created_by         int       `json:"Created_by"`
	// Updated_at         time.Time `json:"Updated_at"`
	// Updated_by         int       `json:"Updated_by"`
	// Deleted_at         time.Time `json:"Deleted_at"`
	// Deleted_by         int       `json:"Deleted_by"`
	// Type               int    `json:"Type"`
	// Remark             string `json:"Remark"`
	// Email_verification string `json:"Email_verification"`
	// Domain string
	// Token string `json:"token,omitempty"`
}

func FromService(u user.Domain) UserResp {
	return UserResp{
		ID:           int(u.ID),
		Fullname:     u.Fullname,
		Company_name: u.Company_name,
		Email:        u.Email,
		Phone:        u.Phone,
		Addres:       u.Address,
		Username:     u.Username,
		Branch_id:    u.Branch_id,
		Privilage:    u.Privilage,
		// Password:     u.Password,
		// Status:             u.Status,
		// Last_login:         u.Last_login,
		// Forgot_key:         u.Forgot_key,
		// Created_at:         u.Created_at,
		// Created_by:         u.Created_by,
		// Updated_at:         u.Updated_at,
		// Updated_by:         u.Updated_by,
		// Deleted_at:         u.Deleted_at,
		// Deleted_by:         u.Deleted_by,
		// Type:               u.Type,
		// Remark:             u.Remark,
		// Email_verification: u.Email_verification,
		// Token: u.Token,
	}
}
