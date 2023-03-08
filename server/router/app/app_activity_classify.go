package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkActivityClassifyRouter struct {
}

// InitHkActivityClassifyRouter 初始化 HkActivityClassify 路由信息
func (s *HkActivityClassifyRouter) InitHkActivityClassifyRouter(Router *gin.RouterGroup) {
	appRouter := Router.Group("app")
	hkActivityClassifyRouter := appRouter.Group("hkActivityClassify").Use(middleware.OperationRecord())
	hkActivityClassifyRouterWithoutRecord := appRouter.Group("hkActivityClassify")
	var hkActivityClassifyApi = v1.ApiGroupApp.AppApiGroup.HkActivityClassifyApi
	{
		hkActivityClassifyRouter.POST("createHkActivityClassify", hkActivityClassifyApi.CreateHkActivityClassify)             // 新建HkActivityClassify
		hkActivityClassifyRouter.DELETE("deleteHkActivityClassify", hkActivityClassifyApi.DeleteHkActivityClassify)           // 删除HkActivityClassify
		hkActivityClassifyRouter.DELETE("deleteHkActivityClassifyByIds", hkActivityClassifyApi.DeleteHkActivityClassifyByIds) // 批量删除HkActivityClassify
		hkActivityClassifyRouter.PUT("updateHkActivityClassify", hkActivityClassifyApi.UpdateHkActivityClassify)              // 更新HkActivityClassify
	}
	{
		hkActivityClassifyRouterWithoutRecord.GET("findHkActivityClassify", hkActivityClassifyApi.FindHkActivityClassify)       // 根据ID获取HkActivityClassify
		hkActivityClassifyRouterWithoutRecord.GET("getHkActivityClassifyList", hkActivityClassifyApi.GetHkActivityClassifyList) // 获取HkActivityClassify列表
	}
}
