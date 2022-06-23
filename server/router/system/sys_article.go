package system

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
	"server/middleware"
)

type ArticleRouter struct {}

func (this *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup)  {
	articleRouter := Router.Group("article").Use(middleware.OperationRecord())
	articleApi := v1.ApiGroupApp.SystemApiGroup.SysArticle
	articleRouterWithoutRecord := Router.Group("article")
	{
		articleRouter.POST("addArticle",articleApi.CreateArticle)
		articleRouter.POST("editArticle",articleApi.EditArticle)
		articleRouter.POST("deleteArticle",articleApi.DeleteArticle)
	}
	{
		articleRouterWithoutRecord.POST("listArticle",articleApi.ListArticle)
		articleRouterWithoutRecord.POST("infoArticle",articleApi.InfoArticle)
	}
}
