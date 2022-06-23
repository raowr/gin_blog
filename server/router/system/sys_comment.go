package system

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
	"server/middleware"
)

type CommentRouter struct {}

func (this *CommentRouter) InitCommentRouter(Router *gin.RouterGroup)  {
	commentRouter := Router.Group("comment").Use(middleware.OperationRecord())
	commentApi := v1.ApiGroupApp.SystemApiGroup.SysComment
	commentRouterWithoutRecord := Router.Group("comment")
	{
		commentRouter.POST("showCommentIdsByIds",commentApi.ShowCommentIdsByIds)
		commentRouter.POST("hideCommentIdsByIds",commentApi.HideCommentIdsByIds)
	}
	{
		commentRouterWithoutRecord.POST("listComment",commentApi.ListComment)
	}
}
