package home

import (
	"github.com/gin-gonic/gin"
	"server/home"
)

type About struct {}

func (this *About)InitAboutRouter(Router *gin.RouterGroup) {
	AboutRouter := Router.Group("home")
	HomeGroupAppRouter := home.HomeGroupApp
	Router.GET("/about",HomeGroupAppRouter.AboutGroup.About)
	{
		AboutRouter.GET("about",HomeGroupAppRouter.AboutGroup.About)
	}
}
