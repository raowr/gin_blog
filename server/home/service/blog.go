package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server/global"
	"server/home/msg"
	"server/model/common/request"
	"server/model/common/response"
	"server/model/home"
	"server/utils"
)

type ServiceBlog struct {}

//List article列表
func (this *ServiceBlog)List(article home.Article, info request.PageInfo) (articleList []home.Article,total int64,err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&home.Article{})

	if article.Title != ""{
		db = db.Where("title LIKE ?", "%"+article.Title+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		fmt.Errorf("article列表查询错误: %v", err)
	}else {
		err = db.Limit(limit).Offset(offset).Find(&articleList).Error
	}
	return articleList, total,err
}

func (this *ServiceBlog) InfoArticle(blogID uint) (info home.Article,err error) {
	err = global.GVA_DB.Where("`id` = ?", blogID).First(&info).Error
	return
}

func (this *ServiceBlog) InfoArticleByWhere(where string,value interface{},order string) (info home.Article,err error) {
	err = global.GVA_DB.Select("id").Where(where,value).First(&info).Order(order).Error
	return
}

func (this *ServiceBlog) BrowsesArticle(ID uint) error {
	return global.GVA_DB.Model(&home.Article{}).Where("`id` = ?", ID).Update("browses", gorm.Expr("browses + ?",1)).Error
}

func (this *ServiceBlog) CommentsArticle(ID uint,comments int64) error {
	return global.GVA_DB.Model(&home.Article{}).Where("`id` = ?", ID).Update("comments", comments).Error
}

func (this *ServiceBlog) FindID(c *gin.Context,to string ) (id uint, err error) {
	var blogInfo msg.SearchListBlogParams
	_= c.ShouldBindJSON(&blogInfo)
	if err = utils.Verify(blogInfo, utils.InfoArticle); err != nil {
		response.FailWithMessage(err.Error(), c)
		return 0,err
	}
	var whereStr string
	var orderStr string
	if to == "pre"{
		whereStr = "id < ? "
		orderStr = "id desc"
	}else {
		whereStr = "id > ? "
		orderStr = "id asc"
	}
	var blogDetail home.Article
	if blogDetail, err = this.InfoArticleByWhere(whereStr,blogInfo.ID,orderStr) ; err != nil{
		return blogInfo.ID,err
	}else {
		return blogDetail.ID,err
	}

}
