package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkMiniProgramPacketRouter struct {
}

// InitHkMiniProgramPacketRouter 初始化 HkMiniProgramPacket 路由信息
func (s *HkMiniProgramPacketRouter) InitHkMiniProgramPacketRouter(Router *gin.RouterGroup) {
	hkMiniProgramPacketRouter := Router.Group("hkMiniProgramPacket").Use(middleware.OperationRecord())
	hkMiniProgramPacketRouterWithoutRecord := Router.Group("hkMiniProgramPacket")
	var hkMiniProgramPacketApi = v1.ApiGroupApp.GeneralApiGroup.HkMiniProgramPacketApi
	{
		hkMiniProgramPacketRouter.POST("createHkMiniProgramPacket", hkMiniProgramPacketApi.CreateHkMiniProgramPacket)   // 新建HkMiniProgramPacket
		hkMiniProgramPacketRouter.DELETE("deleteHkMiniProgramPacket", hkMiniProgramPacketApi.DeleteHkMiniProgramPacket) // 删除HkMiniProgramPacket
		hkMiniProgramPacketRouter.DELETE("deleteHkMiniProgramPacketByIds", hkMiniProgramPacketApi.DeleteHkMiniProgramPacketByIds) // 批量删除HkMiniProgramPacket
		hkMiniProgramPacketRouter.PUT("updateHkMiniProgramPacket", hkMiniProgramPacketApi.UpdateHkMiniProgramPacket)    // 更新HkMiniProgramPacket
	}
	{
		hkMiniProgramPacketRouterWithoutRecord.GET("findHkMiniProgramPacket", hkMiniProgramPacketApi.FindHkMiniProgramPacket)        // 根据ID获取HkMiniProgramPacket
		hkMiniProgramPacketRouterWithoutRecord.GET("getHkMiniProgramPacketList", hkMiniProgramPacketApi.GetHkMiniProgramPacketList)  // 获取HkMiniProgramPacket列表
	}
}
