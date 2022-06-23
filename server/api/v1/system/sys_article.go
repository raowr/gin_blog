package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/global"
	"server/model/common/response"
	"server/model/home"
	"server/model/system/request"
	"server/utils"
)

type SysArticle struct {}

func (this *SysArticle) CreateArticle(c *gin.Context)  {
	var articleInfo home.Article
	_ = c.ShouldBindJSON(&articleInfo)
	if err := utils.Verify(articleInfo, utils.Article); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := articleService.CreateArticle(articleInfo); err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Error(err))

		response.FailWithMessage("添加失败", c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

func (this *SysArticle) ListArticle(c *gin.Context)  {
	var searchArticleParams request.SearchArticleParams
	_ = c.ShouldBindJSON(&searchArticleParams)
	if err := utils.Verify(searchArticleParams, utils.ListArticle); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := articleService.ListArticle(searchArticleParams); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     searchArticleParams.Page,
			PageSize: searchArticleParams.PageSize,
		}, "获取成功", c)
	}
}

func (this *SysArticle) InfoArticle(c *gin.Context)  {
	var searchArticleParamsData request.SearchArticleParams
	_ = c.ShouldBindJSON(&searchArticleParamsData)
	if err := utils.Verify(searchArticleParamsData, utils.InfoArticle); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if info , err := articleService.InfoArticle(searchArticleParamsData); err != nil{
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}else {
		response.OkWithDetailed(gin.H{"infoArticle": info}, "获取成功", c)
	}
}

func (this *SysArticle) EditArticle(c *gin.Context)  {
	var editArticleParamsData request.EditArticleParams
	_ = c.ShouldBindJSON(&editArticleParamsData)
	if err := utils.Verify(editArticleParamsData, utils.EditArticle); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := articleService.EditArticle(editArticleParamsData); err != nil{
		global.GVA_LOG.Error("编辑失败!", zap.Error(err))
		response.FailWithMessage("编辑失败", c)
	}else {
		response.OkWithDetailed(gin.H{}, "编辑成功", c)
	}
}

func (this *SysArticle) DeleteArticle(c *gin.Context)  {
	var searchArticleParamsData request.SearchArticleParams
	_ = c.ShouldBindJSON(&searchArticleParamsData)
	if err := utils.Verify(searchArticleParamsData, utils.InfoArticle); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := articleService.DeteteArticle(searchArticleParamsData); err != nil{
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	}else {
		response.OkWithDetailed(gin.H{}, "删除成功", c)
	}
}


