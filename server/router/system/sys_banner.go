package system

import (
	"github.com/gin-gonic/gin"
	v1 "server/api/v1"
	"server/middleware"
)

type BannerRouter struct {}

func (this *BannerRouter) InitBannerRouter(Router *gin.RouterGroup)  {
	bannerRouter := Router.Group("banner").Use(middleware.OperationRecord())
	bannerApi := v1.ApiGroupApp.SystemApiGroup.SysBanner
	bannerRouterWithoutRecord := Router.Group("banner")
	{
		bannerRouter.POST("addBanner",bannerApi.CreateBanner)
		bannerRouter.POST("editBanner",bannerApi.EditBanner)
		bannerRouter.POST("deleteBanner",bannerApi.DeleteBanner)
	}
	{
		bannerRouterWithoutRecord.POST("listBanner",bannerApi.ListBanner)
		bannerRouterWithoutRecord.POST("infoBanner",bannerApi.InfoBanner)
	}
}
