package system

import (
	"server/global"
	"server/model/home"
	"server/model/system/request"
	"time"
)

type ArticleService struct {}

func (this *ArticleService) CreateArticle(ArticleInfo home.Article) error {
	return global.GVA_DB.Create(&ArticleInfo).Error
}

func (this *ArticleService) ListArticle(SearchArticleParams request.SearchArticleParams) (list []home.Article, total int64, err error) {
	limit := SearchArticleParams.PageSize
	offset := SearchArticleParams.PageSize * (SearchArticleParams.Page - 1)
	db := global.GVA_DB.Model(&home.Article{})

	if SearchArticleParams.Title != ""{
		db = db.Where("title LIKE ?","%"+SearchArticleParams.Title+"%")
	}

	nilTime := time.Time{}
	if SearchArticleParams.CreatedAtStart != nilTime{
		db = db.Where("created_at > ?",SearchArticleParams.CreatedAtStart)
	}

	if SearchArticleParams.CreatedAtEnd != nilTime{
		db = db.Where("created_at < ?",SearchArticleParams.CreatedAtEnd)
	}
	err = db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&list).Error
	return list, total, err

}

func (this *ArticleService) InfoArticle(SearchArticleParams request.SearchArticleParams) (info home.Article,err error) {
	err = global.GVA_DB.Where("`id` = ?", SearchArticleParams.ID).First(&info).Error
	return
}

func (this *ArticleService) EditArticle(EditArticleParams request.EditArticleParams) (err error) {
	err = global.GVA_DB.Model(&home.Article{}).Where("id = ?",EditArticleParams.Id).Updates(map[string]interface{}{
		"title": EditArticleParams.Title,
		"cover": EditArticleParams.Cover,
		"content": EditArticleParams.Content,
	}).Error
	return
}

func (this *ArticleService) DeteteArticle(SearchArticleParams request.SearchArticleParams) (err error) {
	err = global.GVA_DB.Where("id = ?",SearchArticleParams.ID).Delete(&home.Article{}).Error
	return
}
