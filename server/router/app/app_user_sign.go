package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkUserSignRouter struct {
}

// InitHkUserSignRouter 初始化 HkUserSign 路由信息
func (s *HkUserSignRouter) InitHkUserSignRouter(Router *gin.RouterGroup) {
	appRouter := Router.Group("app")
	hkUserSignRouter := appRouter.Group("hkUserSign").Use(middleware.OperationRecord())
	hkUserSignRouterWithoutRecord := appRouter.Group("hkUserSign")
	var hkUserSignApi = v1.ApiGroupApp.AppApiGroup.HkUserSignApi
	{
		hkUserSignRouter.POST("createHkUserSign", hkUserSignApi.CreateHkUserSign)             // 新建HkUserSign
		hkUserSignRouter.DELETE("deleteHkUserSign", hkUserSignApi.DeleteHkUserSign)           // 删除HkUserSign
		hkUserSignRouter.DELETE("deleteHkUserSignByIds", hkUserSignApi.DeleteHkUserSignByIds) // 批量删除HkUserSign
		hkUserSignRouter.PUT("updateHkUserSign", hkUserSignApi.UpdateHkUserSign)              // 更新HkUserSign
	}
	{
		hkUserSignRouterWithoutRecord.GET("findHkUserSign", hkUserSignApi.FindHkUserSign)       // 根据ID获取HkUserSign
		hkUserSignRouterWithoutRecord.GET("getHkUserSignList", hkUserSignApi.GetHkUserSignList) // 获取HkUserSign列表
	}
}
