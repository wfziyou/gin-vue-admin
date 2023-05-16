package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkChannelRouter struct {
}

// InitHkChannelRouter 初始化 HkChannel 路由信息
func (s *HkChannelRouter) InitHkChannelRouter(Router *gin.RouterGroup) {
	appRouter := Router.Group("app")
	hkChannelRouter := appRouter.Group("channel").Use(middleware.OperationRecord())
	hkChannelRouterWithoutRecord := appRouter.Group("channel")
	var hkChannelApi = v1.ApiGroupApp.AppApiGroup.HkChannelApi
	{
		hkChannelRouter.POST("setUserChannel", hkChannelApi.SetUserChannel)                           // 设置用户频道
		hkChannelRouterWithoutRecord.GET("getChannelList", hkChannelApi.GetChannelList)               // 获取频道列表
		hkChannelRouterWithoutRecord.GET("getUserChannelList", hkChannelApi.GetUserChannelList)       // 获取用户频道
		hkChannelRouterWithoutRecord.GET("getGeneralChannelList", hkChannelApi.GetGeneralChannelList) // 获取常规频道列表
	}
}
