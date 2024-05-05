package jwt

import (
	"authentication_service/config"
	"fmt"
	"log"
	"net/http"
)

type JwtService interface {
	GenerateJwt(userID int, userPrivilege int) (string, error)
}
type jwtService struct {
}

func NewAuthService() JwtService {
	return &jwtService{}
}

func (c *jwtService) GenerateJwt(userID int, userPrivilege int) (string, error) {
	server := config.GetJwtServer()
	baseUrl := fmt.Sprintf("http://%s:%s/GenerateToken?id=%d&privilege=%d", server.Host, server.Port, userID, userPrivilege)
	req, err := http.NewRequest("GET", baseUrl, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
		return "", err
	}
	defer resp.Body.Close()

	token := resp.Header.Get("Token")
	return token, err
}
