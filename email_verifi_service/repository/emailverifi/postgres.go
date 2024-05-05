package emailverifi

import (
	"email_verifi_service/business/emailverifi"

	"gorm.io/gorm"
)

type emailverifiRepo struct {
	db *gorm.DB
}

func NewemailverifiRepo(db *gorm.DB) emailverifi.EmailVerifiRepo {
	return &emailverifiRepo{
		db: db,
	}
}

func (r *emailverifiRepo) FindEmailVerifi(userID int) (*emailverifi.Domain, error) {
	var record Emailverifi
	res := r.db.Table("users").Where("id = ?", userID).Take(&record)
	if res.Error != nil {
		return nil, res.Error
	}
	data := record.toService()
	return &data, nil
}
