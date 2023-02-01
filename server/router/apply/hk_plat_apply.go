package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkPlatApplyRouter struct {
}

// InitHkPlatApplyRouter 初始化 HkPlatApply 路由信息
func (s *HkPlatApplyRouter) InitHkPlatApplyRouter(Router *gin.RouterGroup) {
	hkPlatApplyRouter := Router.Group("hkPlatApply").Use(middleware.OperationRecord())
	hkPlatApplyRouterWithoutRecord := Router.Group("hkPlatApply")
	var hkPlatApplyApi = v1.ApiGroupApp.ApplyApiGroup.HkPlatApplyApi
	{
		hkPlatApplyRouter.POST("createHkPlatApply", hkPlatApplyApi.CreateHkPlatApply)   // 新建HkPlatApply
		hkPlatApplyRouter.DELETE("deleteHkPlatApply", hkPlatApplyApi.DeleteHkPlatApply) // 删除HkPlatApply
		hkPlatApplyRouter.DELETE("deleteHkPlatApplyByIds", hkPlatApplyApi.DeleteHkPlatApplyByIds) // 批量删除HkPlatApply
		hkPlatApplyRouter.PUT("updateHkPlatApply", hkPlatApplyApi.UpdateHkPlatApply)    // 更新HkPlatApply
	}
	{
		hkPlatApplyRouterWithoutRecord.GET("findHkPlatApply", hkPlatApplyApi.FindHkPlatApply)        // 根据ID获取HkPlatApply
		hkPlatApplyRouterWithoutRecord.GET("getHkPlatApplyList", hkPlatApplyApi.GetHkPlatApplyList)  // 获取HkPlatApply列表
	}
}
