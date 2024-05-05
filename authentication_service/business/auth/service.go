package auth

import (
	"errors"
	"fmt"
	"log"
	"time"

	authorizationService "authentication_service/business/authorization"
	"authentication_service/business/fcmtoken"
	fcmTokenService "authentication_service/business/fcmtoken"
	jwtService "authentication_service/business/jwt"
	userService "authentication_service/business/user"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(user userService.Domain) (*userService.Domain, error)
}

type authService struct {
	userService          userService.UserService
	authorizationService authorizationService.AuthorizationService
	JWTService           jwtService.JwtService
	fcmTokenService      fcmTokenService.FcmTokenService
}

func NewAuthService(
	userService userService.UserService,
	authorizationService authorizationService.AuthorizationService,
	JWTService jwtService.JwtService,
	fcmTokenService fcmTokenService.FcmTokenService,
) AuthService {
	return &authService{
		userService:          userService,
		authorizationService: authorizationService,
		JWTService:           JWTService,
		fcmTokenService:      fcmTokenService,
	}
}

func (c *authService) Login(user userService.Domain) (*userService.Domain, error) {
	err := c.VerifyCredential(user.Username, user.Password)
	if err != nil {
		return nil, errors.New("Invalid username or password")
	}

	usr, err := c.userService.FindByUsername(user.Username)
	if err != nil {
		return nil, err
	}
	// get privilege
	privilage, err := c.authorizationService.FindUserPrivilege(int(usr.ID))
	if err != nil {
		return nil, err
	}

	usr.Privilage = privilage

	// generate token
	token, err := c.JWTService.GenerateJwt(int(usr.ID), privilage)
	if err != nil {
		return nil, err
	}
	usr.Token = token

	// inser fcm
	if user.FcmToken != "" {
		var fcmToken fcmtoken.FcmToken
		fcmToken.Token = user.FcmToken
		fcmToken.Device_id = token
		fcmToken.Username = usr.Username
		fcmToken.Created_at = time.Now()
		fcmToken.Updated_at = time.Now()
		fcm := c.fcmTokenService.InserFcmToken(fcmToken)
		if err != nil {
			fmt.Println(fcm)
		}
		fmt.Println("Insert fcm token:", user.FcmToken)
	} else {
		fmt.Println("fcm nul")
	}
	return usr, nil
}

func (c *authService) VerifyCredential(username string, password string) error {
	user, err := c.userService.FindByUsername(username)
	if err != nil {
		println(err.Error())
		return err
	}

	isValidPassword := comparePassword(user.Password, []byte(password))
	if !isValidPassword {
		return errors.New("failed to login. check your credential")
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
