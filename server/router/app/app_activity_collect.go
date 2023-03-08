package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkActivityCollectRouter struct {
}

// InitHkActivityCollectRouter 初始化 HkActivityCollect 路由信息
func (s *HkActivityCollectRouter) InitHkActivityCollectRouter(Router *gin.RouterGroup) {
	appRouter := Router.Group("app")
	hkActivityCollectRouter := appRouter.Group("hkActivityCollect").Use(middleware.OperationRecord())
	hkActivityCollectRouterWithoutRecord := appRouter.Group("hkActivityCollect")
	var hkActivityCollectApi = v1.ApiGroupApp.AppApiGroup.HkActivityCollectApi
	{
		hkActivityCollectRouter.POST("createHkActivityCollect", hkActivityCollectApi.CreateHkActivityCollect)             // 新建HkActivityCollect
		hkActivityCollectRouter.DELETE("deleteHkActivityCollect", hkActivityCollectApi.DeleteHkActivityCollect)           // 删除HkActivityCollect
		hkActivityCollectRouter.DELETE("deleteHkActivityCollectByIds", hkActivityCollectApi.DeleteHkActivityCollectByIds) // 批量删除HkActivityCollect
		hkActivityCollectRouter.PUT("updateHkActivityCollect", hkActivityCollectApi.UpdateHkActivityCollect)              // 更新HkActivityCollect
	}
	{
		hkActivityCollectRouterWithoutRecord.GET("findHkActivityCollect", hkActivityCollectApi.FindHkActivityCollect)       // 根据ID获取HkActivityCollect
		hkActivityCollectRouterWithoutRecord.GET("getHkActivityCollectList", hkActivityCollectApi.GetHkActivityCollectList) // 获取HkActivityCollect列表
	}
}
