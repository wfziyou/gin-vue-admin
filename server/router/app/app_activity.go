package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ActivityRouter struct {
}

// InitActivityRouter 初始化 Activity 路由信息
func (s *ActivityRouter) InitActivityRouter(Router *gin.RouterGroup) {
	appRouter := Router.Group("app")
	apiRouter := appRouter.Group("activity").Use(middleware.OperationRecord())
	apiRouterWithoutRecord := appRouter.Group("activity")
	var hkActivityApi = v1.ApiGroupApp.AppApiGroup.ActivityApi
	{
		apiRouter.POST("createActivity", hkActivityApi.CreateActivity)               // 创建活动
		apiRouter.DELETE("deleteActivity", hkActivityApi.DeleteActivity)             // 删除活动
		apiRouter.DELETE("deleteActivityByIds", hkActivityApi.DeleteActivityByIds)   // 批量删除活动
		apiRouter.PUT("updateActivity", hkActivityApi.UpdateActivity)                // 更新活动
		apiRouter.POST("joinActivity", hkActivityApi.JoinActivity)                   // 加入活动
		apiRouter.POST("exitActivity", hkActivityApi.ExitActivity)                   // 退出活动
		apiRouter.DELETE("kickOutActivityUsers", hkActivityApi.KickOutActivityUsers) // 踢出活动的用户
	}
	{
		apiRouterWithoutRecord.GET("findActivity", hkActivityApi.FindActivity)                                     // 用id查询活动详情
		apiRouterWithoutRecord.GET("findActivityUser", hkActivityApi.FindActivityUser)                             // 查询活动的用户
		apiRouterWithoutRecord.GET("getActivityUserList", hkActivityApi.GetActivityUserList)                       // 分页获取活动的用户列表
		apiRouterWithoutRecord.GET("getActivityDynamicList", hkActivityApi.GetActivityDynamicList)                 // 分页获取活动的动态列表
		apiRouterWithoutRecord.GET("getGlobalRecommendQuestionList", hkActivityApi.GetCircleRecommendActivityList) // 分页获圈子推荐活动列表
		apiRouterWithoutRecord.GET("getGlobalRecommendActivityList", hkActivityApi.GetGlobalRecommendActivityList) // 分页获全局推荐活动列表
	}
}
