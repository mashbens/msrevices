package changepassword

import (
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type ChangepasswordRepo interface {
	FindByUserID(userID int) (*Domain, error)
	ChangePassword(userID int, newPass string) error
}

type ChangepasswordService interface {
	ChangePassword(userID int, oldPass string, newPass string) (*Domain, error)
	VerifyCredential(userID int, password string) error
	FindByUserID(userID int) (*Domain, error)
}

type changepasswordService struct {
	changepasswordRepo ChangepasswordRepo
}

func NewChangepasswordService(
	changepasswordRepo ChangepasswordRepo,
) ChangepasswordService {
	return &changepasswordService{
		changepasswordRepo: changepasswordRepo,
	}
}

func (c *changepasswordService) ChangePassword(userID int, oldPass string, newPass string) (*Domain, error) {

	u, err := c.changepasswordRepo.FindByUserID(userID)
	if err != nil {
		fmt.Println(err.Error())

		return nil, err
	}
	err = c.VerifyCredential(u.ID, oldPass)
	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("wrong old password")
	}
	hs := hashAndSalt([]byte(newPass))

	err = c.changepasswordRepo.ChangePassword(userID, hs)
	if err != nil {
		fmt.Println(err.Error())

		return nil, err
	}
	return nil, nil
}

func (c *changepasswordService) VerifyCredential(userID int, password string) error {
	user, err := c.changepasswordRepo.FindByUserID(userID)
	if err != nil {
		fmt.Println(err.Error())

		return err
	}

	isValidPassword := comparePassword(user.Password, []byte(password))
	if !isValidPassword {
		return errors.New("wrong old password")
	}
	return nil
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}

func (c *changepasswordService) FindByUserID(userID int) (*Domain, error) {
	user, err := c.changepasswordRepo.FindByUserID(userID)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	return user, nil
}
