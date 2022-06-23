package core

import (
	"fmt"
	"go.uber.org/zap"
	"server/global"
	"server/initialize"
	"server/service/system"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")
	//第一个参数是：在html文件里引用的目录(即写在html的)，第二个参数：存放静态文件的目录，当在html文件中发现以/home/static开头，
	//就会指向./home/view/static*/目录
	Router.Static("/static","./home/view/static")
	//Router.Static("/uploads","./uploads/file")
	//Router.StaticFS("/static", http.Dir("./home/view"))


	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	//time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
