package banner

import (
	"banner_service/business/banner"

	"gorm.io/gorm"
)

type BannerRepo struct {
	db *gorm.DB
}

func NewBannerRepository(db *gorm.DB) banner.BannerRepo {
	return &BannerRepo{
		db: db,
	}
}

func (r *BannerRepo) FindBannerByID(id int) (*banner.Domain, error) {
	var record Home_banner
	res := r.db.Model(&record).Where("id = ? AND status = ?", id, 1).Take(&record)
	if res.Error != nil {
		return nil, res.Error
	}
	data := record.toService()
	return &data, nil
}

func (r *BannerRepo) FindBannerByName(name string) (*banner.Domain, error) {
	var record Home_banner
	res := r.db.Model(&record).Where("name = ? AND status = ?", name, 1).Take(&record)
	if res.Error != nil {
		return nil, res.Error
	}
	data := record.toService()
	return &data, nil
}

func (c *BannerRepo) FindAllBanner(serach string, limit int, offset int) (data []banner.Domain) {
	var record []Home_banner
	res := c.db.Order("id desc").Limit(limit).Offset(offset).Where("status = ?", 1).Find(&record)
	if res.Error != nil {
		return []banner.Domain{}
	}
	return toServiceList(record)
}
func (c *BannerRepo) CountData() int {
	var record Home_banner
	res := c.db.Model(&record).Where("status = ?", 1).Count(&record.Count)
	if res.Error != nil {
		return 0
	}
	return int(record.Count)
}
