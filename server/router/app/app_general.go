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
		generalRouter.POST("createBugReport", generalApi.CreateBugReport) //Bug反馈

	}
	{
		generalRouterWithoutRecord.GET("getProtocolList", generalApi.GetProtocolList)   //分页获取协议列表
		generalRouterWithoutRecord.GET("findBugReport", generalApi.FindBugReport)       //用id查询Bug反馈
		generalRouterWithoutRecord.GET("getBugReportList", generalApi.GetBugReportList) //分页获取Bug反馈列表
		generalRouterWithoutRecord.GET("findMiniProgram", generalApi.FindMiniProgram)   //用id查询小程序
	}
	return generalRouter
}
