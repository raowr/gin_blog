package home

import (
	"github.com/gin-gonic/gin"
	"server/home"
)

type Comment struct {}

func (this *About)InitCommentRouter(Router *gin.RouterGroup) {
	AboutRouter := Router.Group("home")
	HomeGroupAppRouter := home.HomeGroupApp
	{
		AboutRouter.POST("createComment",HomeGroupAppRouter.HandleComment.Create)
		AboutRouter.POST("listComment",HomeGroupAppRouter.HandleComment.List)
	}
}
