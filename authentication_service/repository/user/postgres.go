package user

import (
	"authentication_service/business/user"

	"gorm.io/gorm"
)

type PostgresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) user.UserRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) FindByEmail(email string) (*user.Domain, error) {
	var record Users
	res := r.db.Where("email = ?", email).Take(&record)
	if res.Error != nil {
		return nil, res.Error
	}
	data := record.toService()
	return &data, nil

}

func (r *PostgresRepository) FindByUserID(userID int) (*user.Domain, error) {
	var record Users
	res := r.db.Where("id = ?", userID).Take(&record)
	if res.Error != nil {
		return nil, res.Error
	}
	data := record.toService()
	return &data, nil
}
func (r *PostgresRepository) FindByUsername(username string) (*user.Domain, error) {
	var record Users
	res := r.db.Where("username = ?", username).Take(&record)
	if res.Error != nil {
		return nil, res.Error
	}
	data := record.toService()
	return &data, nil
}
func (r *PostgresRepository) UpdateProfile(user user.Domain) (*user.Domain, error) {
	record := fromService(user)
	res := r.db.Model(&record).Where("id = ?", user.ID).Updates(map[string]interface{}{"username": user.Username, "fullname": user.Fullname, "address": user.Address, "company_name": user.Company_name})
	if res.Error != nil {
		return nil, res.Error
	}
	data := record.toService()

	return &data, nil
}
func (r *PostgresRepository) ChangePassword(userID int, newPass string) error {
	var record Users
	res := r.db.Model(&record).Where("id = ?", userID).Update("password", newPass)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
