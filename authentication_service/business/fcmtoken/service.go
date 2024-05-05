package fcmtoken

type FcmTokenRepo interface {
	InserFcmToken(FcmToken) error
	DeleteFcmToken(oldFcm string) error
}
type FcmTokenService interface {
	InserFcmToken(FcmToken) error
	DeleteFcmToken(oldFcm string) error
}
type fcmTokenService struct {
	fcmRepo FcmTokenRepo
}

func NewFcmService(
	fcmRepo FcmTokenRepo,
) FcmTokenService {
	return &fcmTokenService{
		fcmRepo: fcmRepo,
	}
}

func (c *fcmTokenService) InserFcmToken(data FcmToken) error {
	err := c.fcmRepo.InserFcmToken(data)
	if err != nil {
		return err
	}
	return nil
}

func (c *fcmTokenService) DeleteFcmToken(fcm string) error {
	err := c.fcmRepo.DeleteFcmToken(fcm)
	if err != nil {
		return err
	}
	return nil
}
