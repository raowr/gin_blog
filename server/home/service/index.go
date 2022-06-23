package service

import (
	"fmt"
	"server/global"
	"server/model/common/request"
	"server/model/home"
)

type ServiceIndex struct {}

//ListBanner banner列表
func (this *ServiceIndex)ListBanner(banner home.Banner, info request.PageInfo) (bannerList []home.Banner,total int64,err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&home.Banner{})

	if banner.Title != ""{
		db = db.Where("title LIKE ?", "%"+banner.Title+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		fmt.Errorf("banner列表查询错误: %v", err)
	}else {
		err = db.Limit(limit).Offset(offset).Find(&bannerList).Error
	}
	return bannerList, total,err
}

//ListArticle article列表
func (this *ServiceIndex)ListArticle(article home.Article, info request.PageInfo) (articleList []home.Article,total int64,err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&home.Article{})

	if article.Title != ""{
		db = db.Where("title LIKE ?", "%"+article.Title+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		fmt.Errorf("banner列表查询错误: %v", err)
	}else {
		err = db.Limit(limit).Offset(offset).Find(&articleList).Error
	}
	return articleList, total,err
}
