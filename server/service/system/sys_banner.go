package system

import (
	"server/global"
	"server/model/home"
	"server/model/system/request"
	"time"
)

type BannerService struct {}

func (this *BannerService) CreateBanner(BannerInfo home.Banner) error {
	return global.GVA_DB.Create(&BannerInfo).Error
}

func (this *BannerService) ListBanner(SearchBannerParams request.SearchBannerParams) (list []home.Banner, total int64, err error) {
	limit := SearchBannerParams.PageSize
	offset := SearchBannerParams.PageSize * (SearchBannerParams.Page - 1)
	db := global.GVA_DB.Model(&home.Banner{})

	if SearchBannerParams.Title != ""{
		db = db.Where("title LIKE ?","%"+SearchBannerParams.Title+"%")
	}

	nilTime := time.Time{}
	if SearchBannerParams.CreatedAtStart != nilTime{
		db = db.Where("created_at > ?",SearchBannerParams.CreatedAtStart)
	}

	if SearchBannerParams.CreatedAtEnd != nilTime{
		db = db.Where("created_at < ?",SearchBannerParams.CreatedAtEnd)
	}
	err = db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return list, total, err

}

func (this *BannerService) InfoBanner(SearchBannerParams request.SearchBannerParams) (info home.Banner,err error) {
	err = global.GVA_DB.Where("`id` = ?", SearchBannerParams.ID).First(&info).Error
	return
}

func (this *BannerService) EditBanner(EditBannerParams request.EditBannerParams) (err error) {
	err = global.GVA_DB.Model(&home.Banner{}).Where("id = ?",EditBannerParams.Id).Updates(map[string]interface{}{
		"title": EditBannerParams.Title,
		"cover": EditBannerParams.Cover,
		"content": EditBannerParams.Content,
	}).Error
	return
}

func (this *BannerService) DeteteBanner(SearchBannerParams request.SearchBannerParams) (err error) {
	err = global.GVA_DB.Where("id = ?",SearchBannerParams.ID).Delete(&home.Banner{}).Error
	return
}
