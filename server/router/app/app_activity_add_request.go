package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ActivityAddRequestRouter struct {
}

// InitActivityAddRequestRouter 初始化 ActivityAddRequest 路由信息
func (s *ActivityAddRequestRouter) InitActivityAddRequestRouter(Router *gin.RouterGroup) {
	appRouter := Router.Group("app")
	apiRouter := appRouter.Group("activityAddRequest").Use(middleware.OperationRecord())
	apiRouterWithoutRecord := appRouter.Group("activityAddRequest")
	var api = v1.ApiGroupApp.AppApiGroup.ActivityAddRequestApi
	{
		apiRouter.POST("createActivityAddRequest", api.CreateActivityAddRequest)             // 新建ActivityAddRequest
		apiRouter.DELETE("deleteActivityAddRequest", api.DeleteActivityAddRequest)           // 删除ActivityAddRequest
		apiRouter.DELETE("deleteActivityAddRequestByIds", api.DeleteActivityAddRequestByIds) // 批量删除ActivityAddRequest
		apiRouter.PUT("updateActivityAddRequest", api.UpdateActivityAddRequest)              // 更新ActivityAddRequest
	}
	{
		apiRouterWithoutRecord.GET("findActivityAddRequest", api.FindActivityAddRequest)       // 根据ID获取ActivityAddRequest
		apiRouterWithoutRecord.GET("getActivityAddRequestList", api.GetActivityAddRequestList) // 获取ActivityAddRequest列表
	}
}
