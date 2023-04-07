package app

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct{}

func (s *BaseRouter) InitAuthRouter(Router *gin.RouterGroup) {
	appRouter := Router.Group("app")
	authRouter := appRouter.Group("auth")
	var authApi = v1.ApiGroupApp.AppApiGroup.AuthApi
	{
		authRouter.POST("thirdPlat", authApi.GetThirdPlat)        //获取第三方平台信息
		authRouter.POST("register", authApi.Register)             //用户注册账号
		authRouter.POST("loginPwd", authApi.LoginPwd)             //用户登录(账号密码)
		authRouter.POST("loginTelephone", authApi.LoginTelephone) //用户登录(手机)
		authRouter.POST("loginThird", authApi.LoginThird)         //用户登录(第三方授权)
		authRouter.POST("getCaptcha", authApi.GetSmsVerification) //获取短信验证码
		authRouter.POST("resetPassword", authApi.ResetPassword)   //重置密码
	}
}
