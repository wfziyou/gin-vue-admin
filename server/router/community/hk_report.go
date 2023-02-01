package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkReportRouter struct {
}

// InitHkReportRouter 初始化 HkReport 路由信息
func (s *HkReportRouter) InitHkReportRouter(Router *gin.RouterGroup) {
	hkReportRouter := Router.Group("hkReport").Use(middleware.OperationRecord())
	hkReportRouterWithoutRecord := Router.Group("hkReport")
	var hkReportApi = v1.ApiGroupApp.CommunityApiGroup.HkReportApi
	{
		hkReportRouter.POST("createHkReport", hkReportApi.CreateHkReport)   // 新建HkReport
		hkReportRouter.DELETE("deleteHkReport", hkReportApi.DeleteHkReport) // 删除HkReport
		hkReportRouter.DELETE("deleteHkReportByIds", hkReportApi.DeleteHkReportByIds) // 批量删除HkReport
		hkReportRouter.PUT("updateHkReport", hkReportApi.UpdateHkReport)    // 更新HkReport
	}
	{
		hkReportRouterWithoutRecord.GET("findHkReport", hkReportApi.FindHkReport)        // 根据ID获取HkReport
		hkReportRouterWithoutRecord.GET("getHkReportList", hkReportApi.GetHkReportList)  // 获取HkReport列表
	}
}
