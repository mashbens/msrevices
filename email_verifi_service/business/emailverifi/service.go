package emailverifi

import "fmt"

type EmailVerifiRepo interface {
	FindEmailVerifi(userID int) (*Domain, error)
}

type EmailVerifiService interface {
	FindEmailVerifi(userID int) (res Result, err error)
}

type emailVerifiService struct {
	emailVerifiRepo EmailVerifiRepo
}

func NewEmailVerifiService(
	emailVerifiRepo EmailVerifiRepo,
) EmailVerifiService {
	return &emailVerifiService{
		emailVerifiRepo: emailVerifiRepo,
	}
}

func (c *emailVerifiService) FindEmailVerifi(userID int) (res Result, err error) {
	user, err := c.emailVerifiRepo.FindEmailVerifi(userID)
	if err != nil {
		fmt.Println(err)
		return res, err
	}
	res.Domain = *user
	return res, nil
}
