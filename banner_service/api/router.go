package api

import (
	"banner_service/api/v1/banner"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	Banner *banner.BannerController
}

func RegisterRoutes(e *echo.Echo, controller *Controller) {
	e.GET("/list", controller.Banner.FindAllBanner)
	e.GET("", controller.Banner.FindBannerByName)
	e.GET("/details/:id", controller.Banner.FindBannerByID)

}
