package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppUserRouter struct {
}

// InitUserRouter 初始化 User 路由信息
func (s *AppUserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	appRouter := Router.Group("app")
	userRouter := appRouter.Group("user").Use(middleware.OperationRecord())
	userRouterWithoutRecord := appRouter.Group("user")
	var userApi = v1.ApiGroupApp.AppApiGroup.UserApi
	{
		userRouter.POST("resetPassword", userApi.ResetPassword) //重置密码
		//userRouter.PUT("updateUser", userApi.UpdateUser)        //更新User
	}
	{
		//userRouterWithoutRecord.GET("findUser", userApi.FindUser)       // 根据ID获取User
		userRouterWithoutRecord.GET("getUserList", userApi.GetUserList) // 获取User列表
	}
	return userRouter
}
