package app

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GeneralRouter struct {
}

func (s *GeneralRouter) InitGeneralRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	appRouter := Router.Group("app")
	router := appRouter.Group("general").Use(middleware.OperationRecord())
	withoutRecord := appRouter.Group("general")
	var api = v1.ApiGroupApp.AppApiGroup.GeneralApi
	{
		router.POST("createBugReport", api.CreateBugReport)         //Bug反馈
		withoutRecord.GET("findBugReport", api.FindBugReport)       //用id查询Bug反馈
		withoutRecord.GET("getBugReportList", api.GetBugReportList) //分页获取Bug反馈列表
		withoutRecord.GET("findMiniProgram", api.FindMiniProgram)   //用id查询小程序
	}
	{
		withoutRecord.GET("getFeedbackTypeList", api.GetFeedbackTypeList) // 获取反馈类型列表
		router.POST("createFeedback", api.CreateFeedback)                 // 创建反馈
		router.DELETE("deleteFeedback", api.DeleteFeedback)               // 删除反馈
		router.DELETE("deleteFeedbackByIds", api.DeleteFeedbackByIds)     // 批量删除反馈
		router.PUT("updateFeedback", api.UpdateFeedback)                  // 更新反馈
		withoutRecord.GET("findFeedback", api.FindFeedback)               // 用id查询反馈
		withoutRecord.GET("getFeedbackList", api.GetFeedbackList)         // 分页获取反馈列表
	}
	return router
}
