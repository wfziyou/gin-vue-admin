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
	hkChannelRouter := appRouter.Group("hkChannel").Use(middleware.OperationRecord())
	hkChannelRouterWithoutRecord := appRouter.Group("hkChannel")
	var hkChannelApi = v1.ApiGroupApp.AppApiGroup.HkChannelApi
	{
		hkChannelRouter.POST("createHkChannel", hkChannelApi.CreateHkChannel)             // 新建HkChannel
		hkChannelRouter.DELETE("deleteHkChannel", hkChannelApi.DeleteHkChannel)           // 删除HkChannel
		hkChannelRouter.DELETE("deleteHkChannelByIds", hkChannelApi.DeleteHkChannelByIds) // 批量删除HkChannel
		hkChannelRouter.PUT("updateHkChannel", hkChannelApi.UpdateHkChannel)              // 更新HkChannel
	}
	{
		hkChannelRouterWithoutRecord.GET("findHkChannel", hkChannelApi.FindHkChannel)       // 根据ID获取HkChannel
		hkChannelRouterWithoutRecord.GET("getHkChannelList", hkChannelApi.GetHkChannelList) // 获取HkChannel列表
	}
}
