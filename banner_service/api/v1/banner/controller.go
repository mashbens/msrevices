package banner

import (
	service "banner_service/business/banner"
	"net/http"
	"strconv"

	"banner_service/api/common/obj"
	_response "banner_service/api/common/response"
	jwtService "banner_service/business/jwt"

	"banner_service/api/v1/banner/resp"

	"github.com/labstack/echo/v4"
)

type BannerController struct {
	bannerService service.BannerService
	jwtService    jwtService.JwtService
}

func NewBannerController(
	bannerService service.BannerService,
	jwtService jwtService.JwtService,

) *BannerController {
	return &BannerController{
		bannerService: bannerService,
		jwtService:    jwtService,
	}
}

func (controller *BannerController) FindBannerByName(c echo.Context) error {
	header := c.Request().Header.Get("Authorization")
	validate := controller.jwtService.ValidateToken(header)
	if header == "" {
		response := _response.BuildErrorResponse("Failed to validate token", obj.EmptyObj{})
		return c.JSON(http.StatusUnauthorized, response)
	}
	if validate != true {
		response := _response.BuildErrorResponse("Token expired or invalid", obj.EmptyObj{})
		return c.JSON(http.StatusUnauthorized, response)
	}

	search := c.Request().URL.Query().Get("search")
	banner, err := controller.bannerService.FindBannerByName(search)
	if err != nil {
		response := _response.BuildErrorResponse("Invalid request body.", false)
		return c.JSON(http.StatusBadRequest, response)
	}
	data := resp.FromServiceDomain(*banner)
	_response := _response.BuildSuccsessResponse("Succsess get banner", true, data)
	return c.JSON(http.StatusOK, _response)
}

func (controller *BannerController) FindAllBanner(c echo.Context) error {
	header := c.Request().Header.Get("Authorization")
	validate := controller.jwtService.ValidateToken(header)
	if header == "" {
		response := _response.BuildErrorResponse("Failed to validate token", obj.EmptyObj{})
		return c.JSON(http.StatusUnauthorized, response)
	}
	if validate != true {
		response := _response.BuildErrorResponse("Token expired or invalid", obj.EmptyObj{})
		return c.JSON(http.StatusUnauthorized, response)
	}
	pageID := c.Request().URL.Query().Get("page")
	intPage, _ := strconv.Atoi(pageID)
	if intPage == 0 {
		intPage = 1
	}

	res, err := controller.bannerService.FindAllBanner("", intPage)
	if err != nil {
		response := _response.BuildErrorResponse("Data not found.", false)
		return c.JSON(http.StatusInternalServerError, response)
	}

	data := resp.FromService(res)
	_response := _response.BuildSuccsessResponse("Succsess get all banners ", true, data)
	return c.JSON(http.StatusOK, _response)
}
func (controller *BannerController) FindBannerByID(c echo.Context) error {
	header := c.Request().Header.Get("Authorization")
	validate := controller.jwtService.ValidateToken(header)
	if header == "" {
		response := _response.BuildErrorResponse("Failed to validate token", obj.EmptyObj{})
		return c.JSON(http.StatusUnauthorized, response)
	}
	if validate != true {
		response := _response.BuildErrorResponse("Token expired or invalid", obj.EmptyObj{})
		return c.JSON(http.StatusUnauthorized, response)
	}

	id := c.Param("id")
	intID, _ := strconv.Atoi(id)
	product, err := controller.bannerService.FindBannerByID(intID)
	if err != nil {
		response := _response.BuildErrorResponse("Invalid request body.", false)
		return c.JSON(http.StatusBadRequest, response)
	}

	data := resp.FromServiceDomain(*product)
	_response := _response.BuildSuccsessResponse("Succsess get produk", true, data)
	return c.JSON(http.StatusOK, _response)
}
