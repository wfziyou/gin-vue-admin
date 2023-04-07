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
	hkActivityRouter := appRouter.Group("activity").Use(middleware.OperationRecord())
	hkActivityRouterWithoutRecord := appRouter.Group("activity")
	var hkActivityApi = v1.ApiGroupApp.AppApiGroup.ActivityApi
	{
		hkActivityRouter.POST("createActivity", hkActivityApi.CreateActivity)             // 创建活动
		hkActivityRouter.DELETE("deleteActivity", hkActivityApi.DeleteActivity)           // 删除活动
		hkActivityRouter.DELETE("deleteActivityByIds", hkActivityApi.DeleteActivityByIds) // 批量删除活动
		hkActivityRouter.PUT("updateActivity", hkActivityApi.UpdateActivity)              // 更新活动

		hkActivityRouter.POST("createActivityClassify", hkActivityApi.CreateActivityClassify)             // 新建ActivityClassify
		hkActivityRouter.DELETE("deleteActivityClassify", hkActivityApi.DeleteActivityClassify)           // 删除活动分类
		hkActivityRouter.DELETE("deleteActivityClassifyByIds", hkActivityApi.DeleteActivityClassifyByIds) // 批量删除活动分类
		hkActivityRouter.PUT("updateActivityClassify", hkActivityApi.UpdateActivityClassify)              // 更新活动分类

		hkActivityRouter.POST("joinActivity", hkActivityApi.JoinActivity)                   // 加入活动
		hkActivityRouter.POST("exitActivity", hkActivityApi.ExitActivity)                   // 退出活动
		hkActivityRouter.DELETE("kickOutActivityUsers", hkActivityApi.KickOutActivityUsers) // 踢出活动的用户
	}
	{
		hkActivityRouterWithoutRecord.GET("findActivity", hkActivityApi.FindActivity)                       // 用id查询活动详情
		hkActivityRouterWithoutRecord.GET("getActivityList", hkActivityApi.GetActivityList)                 // 分页获取推荐活动列表
		hkActivityRouterWithoutRecord.GET("findActivityClassify", hkActivityApi.FindActivityClassify)       // 用id查询活动分类
		hkActivityRouterWithoutRecord.GET("getActivityClassifyList", hkActivityApi.GetActivityClassifyList) // 分页获取活动分类列表
		hkActivityRouterWithoutRecord.GET("findActivityUser", hkActivityApi.FindActivityUser)               // 用id查询活动的用户
		hkActivityRouterWithoutRecord.GET("getActivityUserList", hkActivityApi.GetActivityUserList)         // 分页获取活动的用户列表
	}
}
