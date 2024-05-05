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
	Driver     string
	DB_Host    string
	DB_Port    string
	DB_User    string
	DB_Pass    string
	DB_Name    string
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
		Server_Url: os.Getenv("Server_url"),
		Driver:     os.Getenv("Driver"),
		DB_Host:    os.Getenv("Db_host"),
		DB_Port:    os.Getenv("Db_port"),
		DB_User:    os.Getenv("Db_user"),
		DB_Pass:    os.Getenv("Db_pass"),
		DB_Name:    os.Getenv("Db_name"),
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
	sName, _ := parser.QueryToString("data.[24].serviceName")
	host, _ := parser.QueryToString("data.[24].serviceHost")
	sPort, _ := parser.QueryToString("data.[24].servicePort")
	port := fmt.Sprintf(":%s", sPort)
	fmt.Println("u run " + sName)
	fmt.Println("host = ", host)
	fmt.Println("port = ", sPort)
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
	port, _ := parser.QueryToString("data.[3].servicePort")

	serverConfig.Port = port
	serverConfig.Host = host
	return serverConfig
}
