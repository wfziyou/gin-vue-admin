package app

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ReportRouter struct {
}

func (router *ReportRouter) InitReportRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	appRouter := Router.Group("app")
	reportRouter := appRouter.Group("report").Use(middleware.OperationRecord())
	reportRouterWithoutRecord := appRouter.Group("report")
	var reportApi = v1.ApiGroupApp.AppApiGroup.ReportApi
	{
		reportRouter.POST("createReport", reportApi.CreateReport) //举报
	}
	{
		reportRouterWithoutRecord.GET("findReport", reportApi.FindReport)                   //用id查询Report
		reportRouterWithoutRecord.GET("getReportList", reportApi.GetReportList)             //分页获取Report列表
		reportRouterWithoutRecord.GET("findReportReason", reportApi.FindReportReason)       //用id查询ReportReason
		reportRouterWithoutRecord.GET("getReportReasonList", reportApi.GetReportReasonList) //分页获取ReportReason列表
	}
	return reportRouter
}