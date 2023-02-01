package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkUserCircleApplyRouter struct {
}

// InitHkUserCircleApplyRouter 初始化 HkUserCircleApply 路由信息
func (s *HkUserCircleApplyRouter) InitHkUserCircleApplyRouter(Router *gin.RouterGroup) {
	hkUserCircleApplyRouter := Router.Group("hkUserCircleApply").Use(middleware.OperationRecord())
	hkUserCircleApplyRouterWithoutRecord := Router.Group("hkUserCircleApply")
	var hkUserCircleApplyApi = v1.ApiGroupApp.CommunityApiGroup.HkUserCircleApplyApi
	{
		hkUserCircleApplyRouter.POST("createHkUserCircleApply", hkUserCircleApplyApi.CreateHkUserCircleApply)   // 新建HkUserCircleApply
		hkUserCircleApplyRouter.DELETE("deleteHkUserCircleApply", hkUserCircleApplyApi.DeleteHkUserCircleApply) // 删除HkUserCircleApply
		hkUserCircleApplyRouter.DELETE("deleteHkUserCircleApplyByIds", hkUserCircleApplyApi.DeleteHkUserCircleApplyByIds) // 批量删除HkUserCircleApply
		hkUserCircleApplyRouter.PUT("updateHkUserCircleApply", hkUserCircleApplyApi.UpdateHkUserCircleApply)    // 更新HkUserCircleApply
	}
	{
		hkUserCircleApplyRouterWithoutRecord.GET("findHkUserCircleApply", hkUserCircleApplyApi.FindHkUserCircleApply)        // 根据ID获取HkUserCircleApply
		hkUserCircleApplyRouterWithoutRecord.GET("getHkUserCircleApplyList", hkUserCircleApplyApi.GetHkUserCircleApplyList)  // 获取HkUserCircleApply列表
	}
}
