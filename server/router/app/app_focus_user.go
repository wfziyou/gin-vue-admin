package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkFocusUserRouter struct {
}

// InitHkFocusUserRouter 初始化 HkFocusUser 路由信息
func (s *HkFocusUserRouter) InitHkFocusUserRouter(Router *gin.RouterGroup) {
	appRouter := Router.Group("app")
	hkFocusUserRouter := appRouter.Group("hkFocusUser").Use(middleware.OperationRecord())
	hkFocusUserRouterWithoutRecord := appRouter.Group("hkFocusUser")
	var hkFocusUserApi = v1.ApiGroupApp.AppApiGroup.HkFocusUserApi
	{
		hkFocusUserRouter.POST("createHkFocusUser", hkFocusUserApi.CreateHkFocusUser)             // 新建HkFocusUser
		hkFocusUserRouter.DELETE("deleteHkFocusUser", hkFocusUserApi.DeleteHkFocusUser)           // 删除HkFocusUser
		hkFocusUserRouter.DELETE("deleteHkFocusUserByIds", hkFocusUserApi.DeleteHkFocusUserByIds) // 批量删除HkFocusUser
		hkFocusUserRouter.PUT("updateHkFocusUser", hkFocusUserApi.UpdateHkFocusUser)              // 更新HkFocusUser
	}
	{
		hkFocusUserRouterWithoutRecord.GET("findHkFocusUser", hkFocusUserApi.FindHkFocusUser)       // 根据ID获取HkFocusUser
		hkFocusUserRouterWithoutRecord.GET("getHkFocusUserList", hkFocusUserApi.GetHkFocusUserList) // 获取HkFocusUser列表
	}
}
