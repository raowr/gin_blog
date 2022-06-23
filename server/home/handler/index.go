package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"server/global"
	"server/home/msg"
	"server/model/common/response"
	"server/model/home"
	"server/utils"
)

type HandlerIndex struct {}

func (this *HandlerIndex)List(c *gin.Context)  {
	var pageInfoBanner msg.SearchListBannerParams
	pageInfoBanner.Page = 1
	pageInfoBanner.PageSize = 4
	if err := utils.Verify(pageInfoBanner.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var (
		listBanner []home.Banner
		listArticle []home.Article
		totalBanner,totalArticle int64
		err error
	)
	if listBanner, totalBanner, err = serviceIndex.ListBanner(pageInfoBanner.Banner,pageInfoBanner.PageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}

	var pageInfoArticle msg.SearchListBlogParams
	pageInfoArticle.Page = 1
	pageInfoArticle.PageSize = 5
	if err := utils.Verify(pageInfoArticle.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if listArticle, totalArticle, err = serviceIndex.ListArticle(pageInfoArticle.Article,pageInfoArticle.PageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "首页",//
		"listBanner":     listBanner,
		"totalBanner":    totalBanner,
		"listArticle":     listArticle,
		"totalArticle":    totalArticle,
	})
}
