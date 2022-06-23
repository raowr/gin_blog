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

type SysBanner struct {}

func (this *SysBanner) CreateBanner(c *gin.Context)  {
	var bannerInfo home.Banner
	_ = c.ShouldBindJSON(&bannerInfo)
	if err := utils.Verify(bannerInfo, utils.Banner); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := bannerService.CreateBanner(bannerInfo); err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Error(err))

		response.FailWithMessage("添加失败", c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

func (this *SysBanner) ListBanner(c *gin.Context)  {
	var searchBannerParams request.SearchBannerParams
	_ = c.ShouldBindJSON(&searchBannerParams)
	if err := utils.Verify(searchBannerParams, utils.ListBanner); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := bannerService.ListBanner(searchBannerParams); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     searchBannerParams.Page,
			PageSize: searchBannerParams.PageSize,
		}, "获取成功", c)
	}
}

func (this *SysBanner) InfoBanner(c *gin.Context)  {
	var searchBannerParamsData request.SearchBannerParams
	_ = c.ShouldBindJSON(&searchBannerParamsData)
	if err := utils.Verify(searchBannerParamsData, utils.InfoBanner); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if info , err := bannerService.InfoBanner(searchBannerParamsData); err != nil{
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}else {
		response.OkWithDetailed(gin.H{"infoBanner": info}, "获取成功", c)
	}
}

func (this *SysBanner) EditBanner(c *gin.Context)  {
	var editBannerParamsData request.EditBannerParams
	_ = c.ShouldBindJSON(&editBannerParamsData)
	if err := utils.Verify(editBannerParamsData, utils.EditBanner); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := bannerService.EditBanner(editBannerParamsData); err != nil{
		global.GVA_LOG.Error("编辑失败!", zap.Error(err))
		response.FailWithMessage("编辑失败", c)
	}else {
		response.OkWithDetailed(gin.H{}, "编辑成功", c)
	}
}

func (this *SysBanner) DeleteBanner(c *gin.Context)  {
	var searchBannerParamsData request.SearchBannerParams
	_ = c.ShouldBindJSON(&searchBannerParamsData)
	if err := utils.Verify(searchBannerParamsData, utils.InfoBanner); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := bannerService.DeteteBanner(searchBannerParamsData); err != nil{
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	}else {
		response.OkWithDetailed(gin.H{}, "删除成功", c)
	}
}


