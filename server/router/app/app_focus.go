package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type FocusUserRouter struct {
}

// InitFocusUserRouter 初始化 FocusUser 路由信息
func (s *FocusUserRouter) InitFocusUserRouter(Router *gin.RouterGroup) {
	appRouter := Router.Group("app")
	hkFocusUserRouter := appRouter.Group("focus").Use(middleware.OperationRecord())
	hkFocusUserRouterWithoutRecord := appRouter.Group("focus")
	var focusApi = v1.ApiGroupApp.AppApiGroup.FocusApi
	{
		hkFocusUserRouter.POST("focusUser", focusApi.FocusUser)             // 关注用户
		hkFocusUserRouter.POST("focusUserCancel", focusApi.FocusUserCancel) // 取消用户关注
		hkFocusUserRouter.PUT("updateFocusUser", focusApi.UpdateFocusUser)  // 更新关注用户
	}
	{
		hkFocusUserRouterWithoutRecord.GET("getFrequentBrowsingUserList", focusApi.GetFrequentBrowsingUserList) // 分页获取经常浏览用户列表
		hkFocusUserRouterWithoutRecord.GET("getFocusUserDynamicList", focusApi.GetFocusUserDynamicList)         // 分页获取关注用户动态列表
		hkFocusUserRouterWithoutRecord.GET("getFocusUserList", focusApi.GetFocusUserList)                       //分页获取关注用户列表
		hkFocusUserRouterWithoutRecord.GET("getFansList", focusApi.GetFansList)                                 //分页获取粉丝列表
		hkFocusUserRouterWithoutRecord.GET("findFocusUser", focusApi.FindFocusUser)                             // 用id查询关注用户
	}
}
