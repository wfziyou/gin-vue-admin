package apply

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkCircleApplyRouter struct {
}

// InitHkCircleApplyRouter 初始化 HkCircleApply 路由信息
func (s *HkCircleApplyRouter) InitHkCircleApplyRouter(Router *gin.RouterGroup) {
	hkCircleApplyRouter := Router.Group("hkCircleApply").Use(middleware.OperationRecord())
	hkCircleApplyRouterWithoutRecord := Router.Group("hkCircleApply")
	var hkCircleApplyApi = v1.ApiGroupApp.ApplyApiGroup.HkCircleApplyApi
	{
		hkCircleApplyRouter.POST("createHkCircleApply", hkCircleApplyApi.CreateHkCircleApply)   // 新建HkCircleApply
		hkCircleApplyRouter.DELETE("deleteHkCircleApply", hkCircleApplyApi.DeleteHkCircleApply) // 删除HkCircleApply
		hkCircleApplyRouter.DELETE("deleteHkCircleApplyByIds", hkCircleApplyApi.DeleteHkCircleApplyByIds) // 批量删除HkCircleApply
		hkCircleApplyRouter.PUT("updateHkCircleApply", hkCircleApplyApi.UpdateHkCircleApply)    // 更新HkCircleApply
	}
	{
		hkCircleApplyRouterWithoutRecord.GET("findHkCircleApply", hkCircleApplyApi.FindHkCircleApply)        // 根据ID获取HkCircleApply
		hkCircleApplyRouterWithoutRecord.GET("getHkCircleApplyList", hkCircleApplyApi.GetHkCircleApplyList)  // 获取HkCircleApply列表
	}
}
