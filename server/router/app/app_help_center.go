package app

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HelpCenterRouter struct {
}

func (s *HelpCenterRouter) InitHelpCenterRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	appRouter := Router.Group("app")
	router := appRouter.Group("helpCenter").Use(middleware.OperationRecord())
	withoutRecord := appRouter.Group("helpCenter")
	var api = v1.ApiGroupApp.AppApiGroup.HelpCenterApi
	{
		withoutRecord.GET("getHelpTypeList", api.GetHelpTypeList)     // (管理员)获取帮助类型列表
		withoutRecord.GET("getMainHelpList", api.GetMainHelpList)     // 获取主页帮助列表
		withoutRecord.GET("findHelp", api.FindHelp)                   // 用id查询帮助
		router.POST("helpThumbsUp", api.HelpThumbsUp)                 // 帮助点赞
		withoutRecord.GET("getHelpList", api.GetHelpList)             // 获取帮助列表
		router.POST("createFeedback", api.CreateFeedback)             // 创建反馈
		router.DELETE("deleteFeedback", api.DeleteFeedback)           // (管理员)删除反馈
		router.DELETE("deleteFeedbackByIds", api.DeleteFeedbackByIds) // (管理员)批量删除反馈
		router.PUT("updateFeedback", api.UpdateFeedback)              // (管理员)更新反馈
		withoutRecord.GET("findFeedback", api.FindFeedback)           // 用id查询反馈
		withoutRecord.GET("getFeedbackList", api.GetFeedbackList)     // 分页获取反馈列表
	}
	return router
}
