package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	commonRequest "server/model/common/request"
	"server/model/common/response"
	"server/model/system/request"
	"server/utils"
)

type SysComment struct {}

func (this *SysComment) ListComment(c *gin.Context)  {
	var searchCommentParams request.SearchCommentParams
	_= c.ShouldBindJSON(&searchCommentParams)
	if err := utils.Verify(searchCommentParams, utils.ListBanner); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := commentService.ListComment(searchCommentParams); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     searchCommentParams.Page,
			PageSize: searchCommentParams.PageSize,
		}, "获取成功", c)
	}
}


func (this *SysComment) ShowCommentIdsByIds(c *gin.Context)  {
	var ids commonRequest.IdsReq
	_ = c.ShouldBindJSON(&ids)
	if err := commentService.ChangCommentStatusIdsByIds(ids,1); err != nil {
		global.GVA_LOG.Error("操作失败!", zap.Error(err))
		response.FailWithMessage("操作失败", c)
	} else {
		response.OkWithMessage("操作成功", c)
	}
}

func (this *SysComment) HideCommentIdsByIds(c *gin.Context)  {
	var ids commonRequest.IdsReq
	_ = c.ShouldBindJSON(&ids)
	if err := commentService.ChangCommentStatusIdsByIds(ids,0); err != nil {
		global.GVA_LOG.Error("操作失败!", zap.Error(err))
		response.FailWithMessage("操作失败", c)
	} else {
		response.OkWithMessage("操作成功", c)
	}
}
