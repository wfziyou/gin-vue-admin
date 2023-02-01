package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkProtocolRouter struct {
}

// InitHkProtocolRouter 初始化 HkProtocol 路由信息
func (s *HkProtocolRouter) InitHkProtocolRouter(Router *gin.RouterGroup) {
	hkProtocolRouter := Router.Group("hkProtocol").Use(middleware.OperationRecord())
	hkProtocolRouterWithoutRecord := Router.Group("hkProtocol")
	var hkProtocolApi = v1.ApiGroupApp.GeneralApiGroup.HkProtocolApi
	{
		hkProtocolRouter.POST("createHkProtocol", hkProtocolApi.CreateHkProtocol)   // 新建HkProtocol
		hkProtocolRouter.DELETE("deleteHkProtocol", hkProtocolApi.DeleteHkProtocol) // 删除HkProtocol
		hkProtocolRouter.DELETE("deleteHkProtocolByIds", hkProtocolApi.DeleteHkProtocolByIds) // 批量删除HkProtocol
		hkProtocolRouter.PUT("updateHkProtocol", hkProtocolApi.UpdateHkProtocol)    // 更新HkProtocol
	}
	{
		hkProtocolRouterWithoutRecord.GET("findHkProtocol", hkProtocolApi.FindHkProtocol)        // 根据ID获取HkProtocol
		hkProtocolRouterWithoutRecord.GET("getHkProtocolList", hkProtocolApi.GetHkProtocolList)  // 获取HkProtocol列表
	}
}
