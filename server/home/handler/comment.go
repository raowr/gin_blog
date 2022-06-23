package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/home/msg"
	"server/model/common/response"
	"server/model/home"
	"server/utils"
)

type HandleComment struct {

}
func (this *HandleComment) List(c *gin.Context) {
	var commentParams msg.SearchListCommentParams
	_ = c.ShouldBindJSON(&commentParams)
	if err := utils.Verify(commentParams, utils.ListComment); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if listComment, total, err := serviceComment.List(commentParams); err != nil {
		global.GVA_LOG.Error("获取评论失败!", zap.Error(err))
		response.FailWithMessage("获取评论失败", c)
	}else {
		response.OkWithDetailed(response.PageResult{
			List:     listComment,
			Total:    total,
			Page:     commentParams.Page,
			PageSize: commentParams.PageSize,
		}, "获取成功", c)
	}
}

func (this *HandleComment) Create(c *gin.Context) {
	var createCommentParams home.Comment
	_ = c.ShouldBindJSON(&createCommentParams)
	if err := utils.Verify(createCommentParams, utils.CreateComment); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if  err := serviceComment.Create(createCommentParams); err != nil {
		global.GVA_LOG.Error("评论失败!", zap.Error(err))
		response.FailWithMessage("评论失败", c)
	}else {
		//获取评论
		var commentParams msg.SearchListCommentParams
		commentParams.BlogId = createCommentParams.BlogId
		commentParams.Page = 1
		commentParams.PageSize = 4
		if listComment, total, err := serviceComment.List(commentParams); err != nil {
			global.GVA_LOG.Error("获取评论失败!", zap.Error(err))
			response.FailWithMessage("获取评论失败", c)
		}else {
			//修改评论数
			_ = serviceBlog.CommentsArticle(createCommentParams.BlogId,total)

			response.OkWithDetailed(response.PageResult{
				List:     listComment,
				Total:    total,
				Page:     commentParams.Page,
				PageSize: commentParams.PageSize,
			}, "获取成功", c)
		}
	}
}
