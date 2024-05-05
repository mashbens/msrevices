package authorization

import (
	"authentication_service/config"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/elgs/gojq"
)

type AuthorizationService interface {
	FindUserPrivilege(userID int) (int, error)
}
type authorization struct {
}

func NewAuthService() AuthorizationService {
	return &authorization{}
}

func (c *authorization) FindUserPrivilege(userID int) (int, error) {
	server := config.GetAuthorizationServer()
	baseUrl := fmt.Sprintf("http://%s:%s/AuthenticationUser/%d", server.Host, server.Port, userID)
	req, err := http.Get(baseUrl)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	sb := string(body)
	parser, err := gojq.NewStringQuery(sb)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	privilage, _ := parser.QueryToInt64("data.Privilege")
	return int(privilage), nil
}
