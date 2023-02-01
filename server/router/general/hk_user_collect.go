package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkUserCollectRouter struct {
}

// InitHkUserCollectRouter 初始化 HkUserCollect 路由信息
func (s *HkUserCollectRouter) InitHkUserCollectRouter(Router *gin.RouterGroup) {
	hkUserCollectRouter := Router.Group("hkUserCollect").Use(middleware.OperationRecord())
	hkUserCollectRouterWithoutRecord := Router.Group("hkUserCollect")
	var hkUserCollectApi = v1.ApiGroupApp.GeneralApiGroup.HkUserCollectApi
	{
		hkUserCollectRouter.POST("createHkUserCollect", hkUserCollectApi.CreateHkUserCollect)   // 新建HkUserCollect
		hkUserCollectRouter.DELETE("deleteHkUserCollect", hkUserCollectApi.DeleteHkUserCollect) // 删除HkUserCollect
		hkUserCollectRouter.DELETE("deleteHkUserCollectByIds", hkUserCollectApi.DeleteHkUserCollectByIds) // 批量删除HkUserCollect
		hkUserCollectRouter.PUT("updateHkUserCollect", hkUserCollectApi.UpdateHkUserCollect)    // 更新HkUserCollect
	}
	{
		hkUserCollectRouterWithoutRecord.GET("findHkUserCollect", hkUserCollectApi.FindHkUserCollect)        // 根据ID获取HkUserCollect
		hkUserCollectRouterWithoutRecord.GET("getHkUserCollectList", hkUserCollectApi.GetHkUserCollectList)  // 获取HkUserCollect列表
	}
}
