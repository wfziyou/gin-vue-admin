package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkActivityUserRouter struct {
}

// InitHkActivityUserRouter 初始化 HkActivityUser 路由信息
func (s *HkActivityUserRouter) InitHkActivityUserRouter(Router *gin.RouterGroup) {
	appRouter := Router.Group("app")
	hkActivityUserRouter := appRouter.Group("hkActivityUser").Use(middleware.OperationRecord())
	hkActivityUserRouterWithoutRecord := appRouter.Group("hkActivityUser")
	var hkActivityUserApi = v1.ApiGroupApp.AppApiGroup.HkActivityUserApi
	{
		hkActivityUserRouter.POST("createHkActivityUser", hkActivityUserApi.CreateHkActivityUser)             // 新建HkActivityUser
		hkActivityUserRouter.DELETE("deleteHkActivityUser", hkActivityUserApi.DeleteHkActivityUser)           // 删除HkActivityUser
		hkActivityUserRouter.DELETE("deleteHkActivityUserByIds", hkActivityUserApi.DeleteHkActivityUserByIds) // 批量删除HkActivityUser
		hkActivityUserRouter.PUT("updateHkActivityUser", hkActivityUserApi.UpdateHkActivityUser)              // 更新HkActivityUser
	}
	{
		hkActivityUserRouterWithoutRecord.GET("findHkActivityUser", hkActivityUserApi.FindHkActivityUser)       // 根据ID获取HkActivityUser
		hkActivityUserRouterWithoutRecord.GET("getHkActivityUserList", hkActivityUserApi.GetHkActivityUserList) // 获取HkActivityUser列表
	}
}
