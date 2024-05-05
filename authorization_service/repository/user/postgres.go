package user

import (
	"authorization_service/business/user"

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

func (r *PostgresRepository) FindUserPrivilege(userID int) (*user.Domain, error) {
	var record Users
	res := r.db.Where("id = ?", userID).Take(&record)
	if res.Error != nil {
		return nil, res.Error
	}
	data := record.toService()
	return &data, nil

}
