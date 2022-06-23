package home

import (
	"github.com/gin-gonic/gin"
	"server/home"
)

type Blog struct {

}

func (this *Blog) InitBlogRouter(Router *gin.RouterGroup)  {
	BlogRouter := Router.Group("home")
	HomeGroupAppRouter := home.HomeGroupApp
	Router.GET("/blog/:page/:pageSize",HomeGroupAppRouter.HandleBlog.List)
	Router.GET("/detail/:id",HomeGroupAppRouter.HandleBlog.Detail)
	{
		BlogRouter.GET("blog/:page/:pageSize",HomeGroupAppRouter.HandleBlog.List)
		BlogRouter.POST("listPost",HomeGroupAppRouter.HandleBlog.ListPost)
		BlogRouter.GET("detail/:id",HomeGroupAppRouter.HandleBlog.Detail)
		BlogRouter.POST("toPre",HomeGroupAppRouter.HandleBlog.ToPre)
		BlogRouter.POST("toNext",HomeGroupAppRouter.HandleBlog.ToNext)
	}
}
