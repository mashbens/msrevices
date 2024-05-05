package fcmtoken

import (
	"authentication_service/business/fcmtoken"

	"gorm.io/gorm"
)

type FcmToken struct {
	db *gorm.DB
}

func NewFcmTokenRepo(db *gorm.DB) fcmtoken.FcmTokenRepo {
	return &FcmToken{
		db: db,
	}
}

func (r *FcmToken) InserFcmToken(data fcmtoken.FcmToken) (err error) {
	res := r.db.Table("telegram.fcm_token").Create(&data)
	if res.Error != nil {
		return nil
	}
	return
}

func (r *FcmToken) DeleteFcmToken(fcm string) (err error) {
	var record Fcm_token
	res := r.db.Table("telegram.fcm_token").Where("device_id = ?", fcm).Delete(&record)
	if res.Error != nil {
		return nil
	}
	return
}
