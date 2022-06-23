package home

import (
	"github.com/gin-gonic/gin"
	"server/home"
)

type Contact struct {}

func (this *Contact) InitContactRouter(Router *gin.RouterGroup)  {
	ContactRouter := Router.Group("home")
	HomeGroupAppRouter := home.HomeGroupApp
	Router.GET("/contact",HomeGroupAppRouter.HandleContact.Index)
	{
		ContactRouter.GET("contact",HomeGroupAppRouter.HandleContact.Index)
	}
}
