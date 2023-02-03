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
		userRouter.POST("loginPwd", userApi.LoginPwd)              //用户登录(账号密码)
		userRouter.POST("loginTelephone", userApi.LoginTelephone)  //用户登录(手机)
		userRouter.POST("loginThird", userApi.LoginThird)          //用户登录(第三方授权)
		userRouter.POST("resetPassword", userApi.ResetPassword)    //重置密码
		userRouter.POST("bindTelephone", userApi.BindTelephone)    //绑定手机
		userRouter.POST("bindEmail", userApi.BindEmail)            //绑定邮箱
		userRouter.PUT("setSelfBaseInfo", userApi.SetSelfBaseInfo) //设置用户基础信息

	}
	{
		userRouterWithoutRecord.GET("getUserBaseInfo", userApi.GetUserBaseInfo) // 用id查询UserBaseInfo
		userRouterWithoutRecord.GET("getUserList", userApi.GetUserList)         // 分页获取User列表
	}
	return userRouter
}
