package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkReportReasonRouter struct {
}

// InitHkReportReasonRouter 初始化 HkReportReason 路由信息
func (s *HkReportReasonRouter) InitHkReportReasonRouter(Router *gin.RouterGroup) {
	hkReportReasonRouter := Router.Group("hkReportReason").Use(middleware.OperationRecord())
	hkReportReasonRouterWithoutRecord := Router.Group("hkReportReason")
	var hkReportReasonApi = v1.ApiGroupApp.CommunityApiGroup.HkReportReasonApi
	{
		hkReportReasonRouter.POST("createHkReportReason", hkReportReasonApi.CreateHkReportReason)   // 新建HkReportReason
		hkReportReasonRouter.DELETE("deleteHkReportReason", hkReportReasonApi.DeleteHkReportReason) // 删除HkReportReason
		hkReportReasonRouter.DELETE("deleteHkReportReasonByIds", hkReportReasonApi.DeleteHkReportReasonByIds) // 批量删除HkReportReason
		hkReportReasonRouter.PUT("updateHkReportReason", hkReportReasonApi.UpdateHkReportReason)    // 更新HkReportReason
	}
	{
		hkReportReasonRouterWithoutRecord.GET("findHkReportReason", hkReportReasonApi.FindHkReportReason)        // 根据ID获取HkReportReason
		hkReportReasonRouterWithoutRecord.GET("getHkReportReasonList", hkReportReasonApi.GetHkReportReasonList)  // 获取HkReportReason列表
	}
}
