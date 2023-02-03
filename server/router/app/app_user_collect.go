package app

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserCollectRouter struct {
}

func (router *UserCollectRouter) InitUserCollectRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	appRouter := Router.Group("app")
	userCollectRouter := appRouter.Group("userCollect").Use(middleware.OperationRecord())
	userCollectRouterWithoutRecord := appRouter.Group("userCollect")
	var userCollectApi = v1.ApiGroupApp.AppApiGroup.UserCollectApi
	{
		userCollectRouter.POST("createUserCollect", userCollectApi.CreateUserCollect)             //收藏帖子
		userCollectRouter.DELETE("deleteUserCollect", userCollectApi.DeleteUserCollect)           //删除UserCollect
		userCollectRouter.DELETE("deleteUserCollectByIds", userCollectApi.DeleteUserCollectByIds) //批量删除UserCollect
	}
	{
		userCollectRouterWithoutRecord.GET("getUserCollectList", userCollectApi.GetUserCollectList) //分页获取UserCollect列表

	}
	return userCollectRouter
}
