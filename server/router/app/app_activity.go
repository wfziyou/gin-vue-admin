package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkActivityRouter struct {
}

// InitHkActivityRouter 初始化 HkActivity 路由信息
func (s *HkActivityRouter) InitHkActivityRouter(Router *gin.RouterGroup) {
	appRouter := Router.Group("app")
	hkActivityRouter := appRouter.Group("hkActivity").Use(middleware.OperationRecord())
	hkActivityRouterWithoutRecord := appRouter.Group("hkActivity")
	var hkActivityApi = v1.ApiGroupApp.AppApiGroup.HkActivityApi
	{
		hkActivityRouter.POST("createHkActivity", hkActivityApi.CreateHkActivity)             // 新建HkActivity
		hkActivityRouter.DELETE("deleteHkActivity", hkActivityApi.DeleteHkActivity)           // 删除HkActivity
		hkActivityRouter.DELETE("deleteHkActivityByIds", hkActivityApi.DeleteHkActivityByIds) // 批量删除HkActivity
		hkActivityRouter.PUT("updateHkActivity", hkActivityApi.UpdateHkActivity)              // 更新HkActivity
	}
	{
		hkActivityRouterWithoutRecord.GET("findHkActivity", hkActivityApi.FindHkActivity)       // 根据ID获取HkActivity
		hkActivityRouterWithoutRecord.GET("getHkActivityList", hkActivityApi.GetHkActivityList) // 获取HkActivity列表
	}
}
