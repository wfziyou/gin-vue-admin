package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkPlatApplyGroupRouter struct {
}

// InitHkPlatApplyGroupRouter 初始化 HkPlatApplyGroup 路由信息
func (s *HkPlatApplyGroupRouter) InitHkPlatApplyGroupRouter(Router *gin.RouterGroup) {
	hkPlatApplyGroupRouter := Router.Group("hkPlatApplyGroup").Use(middleware.OperationRecord())
	hkPlatApplyGroupRouterWithoutRecord := Router.Group("hkPlatApplyGroup")
	var hkPlatApplyGroupApi = v1.ApiGroupApp.ApplyApiGroup.HkPlatApplyGroupApi
	{
		hkPlatApplyGroupRouter.POST("createHkPlatApplyGroup", hkPlatApplyGroupApi.CreateHkPlatApplyGroup)   // 新建HkPlatApplyGroup
		hkPlatApplyGroupRouter.DELETE("deleteHkPlatApplyGroup", hkPlatApplyGroupApi.DeleteHkPlatApplyGroup) // 删除HkPlatApplyGroup
		hkPlatApplyGroupRouter.DELETE("deleteHkPlatApplyGroupByIds", hkPlatApplyGroupApi.DeleteHkPlatApplyGroupByIds) // 批量删除HkPlatApplyGroup
		hkPlatApplyGroupRouter.PUT("updateHkPlatApplyGroup", hkPlatApplyGroupApi.UpdateHkPlatApplyGroup)    // 更新HkPlatApplyGroup
	}
	{
		hkPlatApplyGroupRouterWithoutRecord.GET("findHkPlatApplyGroup", hkPlatApplyGroupApi.FindHkPlatApplyGroup)        // 根据ID获取HkPlatApplyGroup
		hkPlatApplyGroupRouterWithoutRecord.GET("getHkPlatApplyGroupList", hkPlatApplyGroupApi.GetHkPlatApplyGroupList)  // 获取HkPlatApplyGroup列表
	}
}
