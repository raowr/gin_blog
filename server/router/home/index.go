package home

import (
	"github.com/gin-gonic/gin"
	"server/home"
)

type Index struct {}

func (this *Index)InitIndexRouter(Router *gin.RouterGroup) {
	IndexRouter := Router.Group("home")
	HomeGroupAppRouter := home.HomeGroupApp
	Router.GET("",HomeGroupAppRouter.IndexGroup.List)
	Router.GET("/index",HomeGroupAppRouter.IndexGroup.List)
	{
		IndexRouter.GET("",HomeGroupAppRouter.IndexGroup.List)
		IndexRouter.GET("index",HomeGroupAppRouter.IndexGroup.List)
	}
}
