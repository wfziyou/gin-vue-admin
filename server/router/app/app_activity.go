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
	var api = v1.ApiGroupApp.AppApiGroup.ActivityApi
	{
		apiRouter.POST("createActivity", api.CreateActivity)                                             // 创建活动
		apiRouter.DELETE("deleteActivity", api.DeleteActivity)                                           // 删除活动
		apiRouter.DELETE("deleteActivityByIds", api.DeleteActivityByIds)                                 // 批量删除活动
		apiRouter.PUT("updateActivity", api.UpdateActivity)                                              // 更新活动
		apiRouter.POST("joinActivity", api.JoinActivity)                                                 // 加入活动
		apiRouter.POST("exitActivity", api.ExitActivity)                                                 // 退出活动
		apiRouter.DELETE("kickOutActivityUsers", api.KickOutActivityUsers)                               // 踢出活动的用户
		apiRouterWithoutRecord.GET("getCircleRecommendActivityList", api.GetCircleRecommendActivityList) // 分页获取圈子推荐活动列表
		apiRouterWithoutRecord.GET("getGlobalRecommendActivityList", api.GetGlobalRecommendActivityList) // 分页获取全局推荐活动列表
		apiRouterWithoutRecord.GET("findActivity", api.FindActivity)                                     // 用id查询活动详情
		apiRouterWithoutRecord.GET("findActivityUser", api.FindActivityUser)                             // 查询活动的用户
		apiRouterWithoutRecord.GET("getActivityUserList", api.GetActivityUserList)                       // 分页获取活动的用户列表

	}
	{
		apiRouter.POST("createActivityDynamic", api.CreateActivityDynamic)               //创建活动动态
		apiRouter.DELETE("deleteActivityDynamic", api.DeleteActivityDynamic)             //删除活动动态
		apiRouter.DELETE("deleteActivityDynamicByIds", api.DeleteActivityDynamicByIds)   //批量删除活动动态
		apiRouterWithoutRecord.GET("getActivityDynamicList", api.GetActivityDynamicList) // 分页获取活动的动态列表
	}
	{
		//活动报名申请
		apiRouter.DELETE("deleteActivityAddRequest", api.DeleteActivityAddRequest)             // 删除活动报名申请
		apiRouter.DELETE("deleteActivityAddRequestByIds", api.DeleteActivityAddRequestByIds)   // 批量删除活动报名申请
		apiRouter.PUT("updateActivityAddRequest", api.UpdateActivityAddRequest)                // 活动报名申请审批
		apiRouterWithoutRecord.GET("findActivityAddRequest", api.FindActivityAddRequest)       // 用id查询活动报名申请
		apiRouterWithoutRecord.GET("getActivityAddRequestList", api.GetActivityAddRequestList) // 分页获取活动报名申请列表
	}
}
