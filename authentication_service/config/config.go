package config

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/elgs/gojq"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type AppConfig struct {
	Server_Url string
	Driver1    string
	DB_Host1   string
	DB_Port1   string
	DB_User1   string
	DB_Pass1   string
	DB_Name1   string

	Driver2  string
	DB_Host2 string
	DB_Port2 string
	DB_User2 string
	DB_Pass2 string
	DB_Name2 string
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}
	return appConfig
}

func initConfig() *AppConfig {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &AppConfig{
		Server_Url: os.Getenv("SERVER"),
		Driver1:    os.Getenv("DRIVER1"),
		DB_Host1:   os.Getenv("DB_HOST1"),
		DB_Port1:   os.Getenv("DB_PORT1"),
		DB_User1:   os.Getenv("DB_USER1"),
		DB_Pass1:   os.Getenv("DB_PASS1"),
		DB_Name1:   os.Getenv("DB_NAME1"),

		Driver2:  os.Getenv("DRIVER2"),
		DB_Host2: os.Getenv("DB_HOST2"),
		DB_Port2: os.Getenv("DB_PORT2"),
		DB_User2: os.Getenv("DB_USER2"),
		DB_Pass2: os.Getenv("DB_PASS2"),
		DB_Name2: os.Getenv("DB_NAME2"),
	}
}

type ServerConfig struct {
	Host string
	Port string
}

func GetServer() ServerConfig {
	var serverConfig ServerConfig
	config := GetConfig()
	prt, err := http.Get(config.Server_Url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(prt.Body)
	if err != nil {
		log.Fatal(err)
	}
	sb := string(body)
	p := strings.TrimSuffix(sb, "\n")
	parser, err := gojq.NewStringQuery(p)
	if err != nil {
		fmt.Println(err)
	}

	sName, _ := parser.QueryToString("data.[1].serviceName")
	host, _ := parser.QueryToString("data.[1].serviceHost")
	sPort, _ := parser.QueryToString("data.[1].servicePort")

	port := fmt.Sprintf(":%s", sPort)
	fmt.Println("u run", sName)
	fmt.Println("host = ", host)
	fmt.Println("port = ", sPort)
	serverConfig.Port = port
	serverConfig.Host = host
	return serverConfig
}
func GetAuthorizationServer() ServerConfig {
	var serverConfig ServerConfig
	config := GetConfig()
	prt, err := http.Get(config.Server_Url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(prt.Body)
	if err != nil {
		log.Fatal(err)
	}
	sb := string(body)
	p := strings.TrimSuffix(sb, "\n")
	parser, err := gojq.NewStringQuery(p)
	if err != nil {
		fmt.Println(err)
	}
	host, _ := parser.QueryToString("data.[2].serviceHost")
	sPort, _ := parser.QueryToString("data.[2].servicePort")
	port := fmt.Sprintf("%s", sPort)

	serverConfig.Port = port
	serverConfig.Host = host
	return serverConfig
}

func GetJwtServer() ServerConfig {
	var serverConfig ServerConfig
	config := GetConfig()
	prt, err := http.Get(config.Server_Url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(prt.Body)
	if err != nil {
		log.Fatal(err)
	}
	sb := string(body)
	p := strings.TrimSuffix(sb, "\n")
	parser, err := gojq.NewStringQuery(p)
	if err != nil {
		fmt.Println(err)
	}

	host, _ := parser.QueryToString("data.[3].serviceHost")
	sPort, _ := parser.QueryToString("data.[3].servicePort")

	port := fmt.Sprintf("%s", sPort)

	serverConfig.Port = port
	serverConfig.Host = host
	return serverConfig
}
