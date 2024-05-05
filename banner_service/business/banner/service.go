package banner

import (
	"math"
)

type BannerRepo interface {
	FindBannerByID(id int) (*Domain, error)
	FindBannerByName(name string) (*Domain, error)
	FindAllBanner(search string, limit int, offset int) (data []Domain)
	CountData() int
}

type BannerService interface {
	FindBannerByID(id int) (*Domain, error)
	FindBannerByName(name string) (*Domain, error)
	FindAllBanner(search string, page int) (res Result, err error)
}

type bannerService struct {
	bannerRepo BannerRepo
}

func NewBannerService(
	bannerRepo BannerRepo,
) BannerService {
	return &bannerService{
		bannerRepo: bannerRepo,
	}
}

func (c *bannerService) FindBannerByID(id int) (*Domain, error) {
	b, err := c.bannerRepo.FindBannerByID(id)
	if err != nil {
		return nil, err
	}
	b.File_mobile = "https://img.gps.id/dev-portal/banner/" + b.File_mobile
	return b, nil
}
func (c *bannerService) FindBannerByName(name string) (*Domain, error) {
	b, err := c.bannerRepo.FindBannerByName(name)
	if err != nil {
		return nil, err
	}
	b.File_mobile = "https://img.gps.id/dev-portal/banner/" + b.File_mobile
	return b, nil
}

func (c *bannerService) FindAllBanner(search string, page int) (res Result, err error) {
	var limit = 10
	offset := page*limit - limit
	if page == 1 {
		offset = 0
	}
	b := c.bannerRepo.FindAllBanner(search, limit, offset)
	domain := []Domain{}
	for _, d := range b {
		d.File_mobile = "https://img.gps.id/dev-portal/banner/" + d.File_mobile
		domain = append(domain, d)
	}

	res.Domain = domain
	count := c.bannerRepo.CountData()
	res.TotalAllData = count
	res.CurrentPage = page
	res.PerPage = limit
	x := float64(count) / float64(limit)
	res.TotalPage = int(math.Ceil(x))
	return res, nil
}
