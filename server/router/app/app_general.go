package app

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GeneralRouter struct {
}

func (router *GeneralRouter) InitGeneralRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	appRouter := Router.Group("app")
	generalRouter := appRouter.Group("general").Use(middleware.OperationRecord())
	generalRouterWithoutRecord := appRouter.Group("general")
	var generalApi = v1.ApiGroupApp.AppApiGroup.GeneralApi
	{
		generalRouter.POST("createBugReport", generalApi.CreateHkBugReport) //Bug反馈

	}
	{
		generalRouterWithoutRecord.GET("findProtocol", generalApi.FindProtocol)         //用id查询协议
		generalRouterWithoutRecord.GET("getProtocolList", generalApi.GetProtocolList)   //分页获取Protocol列表
		generalRouterWithoutRecord.GET("findBugReport", generalApi.FindBugReport)       //用id查询BugReport
		generalRouterWithoutRecord.GET("getBugReportList", generalApi.GetBugReportList) //分页获取BugReport列表
	}
	return generalRouter
}
