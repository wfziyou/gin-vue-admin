package app

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	appRouter := Router.Group("app")
	userRouter := appRouter.Group("user")
	var userApi = v1.ApiGroupApp.AppApiGroup.UserApi
	{
		userRouter.POST("register", userApi.Register)             //用户注册账号
		userRouter.POST("loginPwd", userApi.LoginPwd)             //用户登录(账号密码)
		userRouter.POST("loginTelephone", userApi.LoginTelephone) //用户登录(手机)
		userRouter.POST("loginThird", userApi.LoginThird)         //用户登录(第三方授权)
		userRouter.POST("getCaptcha", userApi.GetCaptcha)         //获取验证码
		userRouter.POST("resetPassword", userApi.ResetPassword)   //重置密码
	}

	generalRouterWithoutRecord := appRouter.Group("general")
	var generalApi = v1.ApiGroupApp.AppApiGroup.GeneralApi
	{
		generalRouterWithoutRecord.GET("findProtocol", generalApi.FindProtocol)             //用id查询协议
		generalRouterWithoutRecord.GET("findProtocolByName", generalApi.FindProtocolByName) //用名字查询协议
	}
}
