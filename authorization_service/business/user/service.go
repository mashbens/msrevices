package user

type UserRepository interface {
	FindUserPrivilege(userID int) (*Domain, error)
}
type UserService interface {
	FindUserPrivilege(userID int) (*Domain, error)
}

type userService struct {
	userRepo UserRepository
}

func NewUserService(
	userRepo UserRepository,

) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (c *userService) FindUserPrivilege(userID int) (*Domain, error) {
	user, err := c.userRepo.FindUserPrivilege(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
