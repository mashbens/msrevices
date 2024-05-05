package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"authentication_service/app/modules"
	"authentication_service/config"
	"authentication_service/util"

	"authentication_service/api"
)

func main() {

	conf := config.GetConfig()
	dbCon1 := util.NewConnectionDatabase(conf)
	dbCon2 := util.NewConnectionDatabase2(conf)
	controllers := modules.RegisterModules(dbCon1, dbCon2, conf)
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	api.RegisterRoutes(e, &controllers)

	server := config.GetServer()

	e.Logger.Fatal(e.Start(server.Port))
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	defer dbCon1.CloseConnection()
	defer dbCon2.CloseConnection2()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
