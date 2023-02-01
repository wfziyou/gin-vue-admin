package general

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkBugReportRouter struct {
}

// InitHkBugReportRouter 初始化 HkBugReport 路由信息
func (s *HkBugReportRouter) InitHkBugReportRouter(Router *gin.RouterGroup) {
	hkBugReportRouter := Router.Group("hkBugReport").Use(middleware.OperationRecord())
	hkBugReportRouterWithoutRecord := Router.Group("hkBugReport")
	var hkBugReportApi = v1.ApiGroupApp.GeneralApiGroup.HkBugReportApi
	{
		hkBugReportRouter.POST("createHkBugReport", hkBugReportApi.CreateHkBugReport)   // 新建HkBugReport
		hkBugReportRouter.DELETE("deleteHkBugReport", hkBugReportApi.DeleteHkBugReport) // 删除HkBugReport
		hkBugReportRouter.DELETE("deleteHkBugReportByIds", hkBugReportApi.DeleteHkBugReportByIds) // 批量删除HkBugReport
		hkBugReportRouter.PUT("updateHkBugReport", hkBugReportApi.UpdateHkBugReport)    // 更新HkBugReport
	}
	{
		hkBugReportRouterWithoutRecord.GET("findHkBugReport", hkBugReportApi.FindHkBugReport)        // 根据ID获取HkBugReport
		hkBugReportRouterWithoutRecord.GET("getHkBugReportList", hkBugReportApi.GetHkBugReportList)  // 获取HkBugReport列表
	}
}
