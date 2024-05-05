package changepassword

import (
	"change_password_service/business/changepassword"
	"fmt"

	"gorm.io/gorm"
)

type changePassword struct {
	db *gorm.DB
}

func NewchangePassword(db *gorm.DB) changepassword.ChangepasswordRepo {
	return &changePassword{
		db: db,
	}
}

func (r *changePassword) FindByUserID(userID int) (*changepassword.Domain, error) {
	var record Data
	res := r.db.Table("users").Where("id = ?", userID).Take(&record)
	if res.Error != nil {
		fmt.Println(res.Error)
		return nil, res.Error
	}
	dataResult := record.toService()
	return &dataResult, nil

}
func (r *changePassword) ChangePassword(userID int, newPass string) error {
	res := r.db.Table("users").Where("id = ?", userID).Update("password", newPass)
	if res.Error != nil {
		fmt.Println(res.Error)
		return res.Error
	}

	return nil

}
