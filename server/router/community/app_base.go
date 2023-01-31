package community

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	appRouter := Router.Group("app")
	userRouter := appRouter.Group("user")
	var userApi = v1.ApiGroupApp.CommunityApiGroup.UserApi
	{
		userRouter.POST("login", userApi.Login)       //用户登录
		userRouter.POST("register", userApi.Register) //用户注册账号
	}
}
