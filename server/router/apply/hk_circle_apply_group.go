package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkCircleApplyGroupRouter struct {
}

// InitHkCircleApplyGroupRouter 初始化 HkCircleApplyGroup 路由信息
func (s *HkCircleApplyGroupRouter) InitHkCircleApplyGroupRouter(Router *gin.RouterGroup) {
	hkCircleApplyGroupRouter := Router.Group("hkCircleApplyGroup").Use(middleware.OperationRecord())
	hkCircleApplyGroupRouterWithoutRecord := Router.Group("hkCircleApplyGroup")
	var hkCircleApplyGroupApi = v1.ApiGroupApp.ApplyApiGroup.HkCircleApplyGroupApi
	{
		hkCircleApplyGroupRouter.POST("createHkCircleApplyGroup", hkCircleApplyGroupApi.CreateHkCircleApplyGroup)   // 新建HkCircleApplyGroup
		hkCircleApplyGroupRouter.DELETE("deleteHkCircleApplyGroup", hkCircleApplyGroupApi.DeleteHkCircleApplyGroup) // 删除HkCircleApplyGroup
		hkCircleApplyGroupRouter.DELETE("deleteHkCircleApplyGroupByIds", hkCircleApplyGroupApi.DeleteHkCircleApplyGroupByIds) // 批量删除HkCircleApplyGroup
		hkCircleApplyGroupRouter.PUT("updateHkCircleApplyGroup", hkCircleApplyGroupApi.UpdateHkCircleApplyGroup)    // 更新HkCircleApplyGroup
	}
	{
		hkCircleApplyGroupRouterWithoutRecord.GET("findHkCircleApplyGroup", hkCircleApplyGroupApi.FindHkCircleApplyGroup)        // 根据ID获取HkCircleApplyGroup
		hkCircleApplyGroupRouterWithoutRecord.GET("getHkCircleApplyGroupList", hkCircleApplyGroupApi.GetHkCircleApplyGroupList)  // 获取HkCircleApplyGroup列表
	}
}
