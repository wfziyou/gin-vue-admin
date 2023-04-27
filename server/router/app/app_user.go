package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

// InitUserRouter 初始化 User 路由信息
func (router *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	appRouter := Router.Group("app")
	userRouter := appRouter.Group("user").Use(middleware.OperationRecord())
	userRouterWithoutRecord := appRouter.Group("user")
	var userApi = v1.ApiGroupApp.AppApiGroup.UserApi
	{
		userRouter.POST("bindTelephone", userApi.BindTelephone)                 //绑定手机
		userRouter.POST("bindEmail", userApi.BindEmail)                         //绑定邮箱
		userRouterWithoutRecord.GET("getUserBaseInfo", userApi.GetUserBaseInfo) // 用id查询用户基本信息
		userRouterWithoutRecord.GET("getUserInfo", userApi.GetUserInfo)         // 用id查询用户信息

		userRouter.PUT("setSelfBaseInfo", userApi.SetSelfBaseInfo)      //设置用户基础信息
		userRouterWithoutRecord.GET("getUserList", userApi.GetUserList) // 分页获取用户列表
	}

	return userRouter
}
