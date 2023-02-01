package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkUserExtendRouter struct {
}

// InitHkUserExtendRouter 初始化 HkUserExtend 路由信息
func (s *HkUserExtendRouter) InitHkUserExtendRouter(Router *gin.RouterGroup) {
	hkUserExtendRouter := Router.Group("hkUserExtend").Use(middleware.OperationRecord())
	hkUserExtendRouterWithoutRecord := Router.Group("hkUserExtend")
	var hkUserExtendApi = v1.ApiGroupApp.CommunityApiGroup.HkUserExtendApi
	{
		hkUserExtendRouter.POST("createHkUserExtend", hkUserExtendApi.CreateHkUserExtend)   // 新建HkUserExtend
		hkUserExtendRouter.DELETE("deleteHkUserExtend", hkUserExtendApi.DeleteHkUserExtend) // 删除HkUserExtend
		hkUserExtendRouter.DELETE("deleteHkUserExtendByIds", hkUserExtendApi.DeleteHkUserExtendByIds) // 批量删除HkUserExtend
		hkUserExtendRouter.PUT("updateHkUserExtend", hkUserExtendApi.UpdateHkUserExtend)    // 更新HkUserExtend
	}
	{
		hkUserExtendRouterWithoutRecord.GET("findHkUserExtend", hkUserExtendApi.FindHkUserExtend)        // 根据ID获取HkUserExtend
		hkUserExtendRouterWithoutRecord.GET("getHkUserExtendList", hkUserExtendApi.GetHkUserExtendList)  // 获取HkUserExtend列表
	}
}
