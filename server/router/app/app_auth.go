package app

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct{}

func (s *BaseRouter) InitAuthRouter(PublicRouter *gin.RouterGroup, PrivateRouter *gin.RouterGroup) {
	authRouterPublic := PublicRouter.Group("app").Group("auth")
	authRouterPrivate := PublicRouter.Group("app").Group("auth")
	var authApi = v1.ApiGroupApp.AppApiGroup.AuthApi
	{
		authRouterPublic.POST("getThirdPlat", authApi.GetThirdPlat)                            //获取第三方平台信息
		authRouterPublic.POST("register", authApi.Register)                                    //用户注册账号
		authRouterPublic.POST("loginPwd", authApi.LoginPwd)                                    //用户登录(账号密码)
		authRouterPublic.POST("loginTelephone", authApi.LoginTelephone)                        //用户登录(手机)
		authRouterPublic.POST("loginThird", authApi.LoginThird)                                //用户登录(第三方授权)
		authRouterPublic.POST("loginOneClick", authApi.LoginOneClick)                          //一键登录
		authRouterPublic.POST("getSmsVerification", authApi.GetSmsVerification)                //获取短信验证码
		authRouterPublic.POST("sendEmailVerification", authApi.SendEmailVerification)          //发送Email验证码
		authRouterPublic.POST("resetPassword", authApi.ResetPassword)                          //重置密码
		authRouterPublic.POST("getLocalTelephone", authApi.GetLocalTelephone)                  //获取本机手机号码
		authRouterPrivate.POST("getSmsVerificationPrivate", authApi.GetSmsVerificationPrivate) //获取短信验证码
	}
}
