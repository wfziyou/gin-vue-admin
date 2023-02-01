package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkApplyRouter struct {
}

// InitHkApplyRouter 初始化 HkApply 路由信息
func (s *HkApplyRouter) InitHkApplyRouter(Router *gin.RouterGroup) {
	hkApplyRouter := Router.Group("hkApply").Use(middleware.OperationRecord())
	hkApplyRouterWithoutRecord := Router.Group("hkApply")
	var hkApplyApi = v1.ApiGroupApp.ApplyApiGroup.HkApplyApi
	{
		hkApplyRouter.POST("createHkApply", hkApplyApi.CreateHkApply)   // 新建HkApply
		hkApplyRouter.DELETE("deleteHkApply", hkApplyApi.DeleteHkApply) // 删除HkApply
		hkApplyRouter.DELETE("deleteHkApplyByIds", hkApplyApi.DeleteHkApplyByIds) // 批量删除HkApply
		hkApplyRouter.PUT("updateHkApply", hkApplyApi.UpdateHkApply)    // 更新HkApply
	}
	{
		hkApplyRouterWithoutRecord.GET("findHkApply", hkApplyApi.FindHkApply)        // 根据ID获取HkApply
		hkApplyRouterWithoutRecord.GET("getHkApplyList", hkApplyApi.GetHkApplyList)  // 获取HkApply列表
	}
}
