package user

type UserRepository interface {
	FindByUsername(username string) (*Domain, error)
	FindByEmail(email string) (*Domain, error)
	FindByUserID(userID int) (*Domain, error)
}
type UserService interface {
	FindByEmail(email string) (*Domain, error)
	FindByUsername(username string) (*Domain, error)
	FindByUserID(userID int) (*Domain, error)
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

func (c *userService) FindByEmail(email string) (*Domain, error) {
	usr, err := c.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (c *userService) FindByUserID(userID int) (*Domain, error) {
	user, err := c.userRepo.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (c *userService) FindByUsername(username string) (*Domain, error) {
	user, err := c.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	return user, nil
}
