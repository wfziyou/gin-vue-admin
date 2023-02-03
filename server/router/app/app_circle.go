package app

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CircleRouter struct {
}

func (router *CircleRouter) InitCircleRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	appRouter := Router.Group("app")
	circleRouter := appRouter.Group("circle").Use(middleware.OperationRecord())
	circleRouterWithoutRecord := appRouter.Group("circle")
	var circleApi = v1.ApiGroupApp.AppApiGroup.CircleApi
	{
		circleRouter.POST("updateCircle", circleApi.UpdateCircle)                     //(圈子管理者)更新Circle
		circleRouter.POST("setUserCurCircle", circleApi.SetUserCurCircle)             //设置用户当前圈子
		circleRouter.DELETE("deleteHkCircleUser", circleApi.DeleteCircleUser)         //删除CircleUser
		circleRouter.DELETE("deleteCircleUserByIds", circleApi.DeleteCircleUserByIds) //批量删除CircleUser
		circleRouter.PUT("updateCircleUser", circleApi.UpdateCircleUser)              //更新CircleUser
	}
	{
		circleRouterWithoutRecord.GET("getCircleForumPostsList", circleApi.GetCircleForumPostsList)         // 分页获取圈子ForumPosts列表
		circleRouterWithoutRecord.GET("getUserCircleForumPostsList", circleApi.GetUserCircleForumPostsList) // 分页获取用户圈子ForumPosts列表
		circleRouterWithoutRecord.GET("getSelfCircleList", circleApi.GetSelfCircleList)                     // 分页获取用户加入的Circle列表
		circleRouterWithoutRecord.GET("findCircle", circleApi.FindCircle)                                   // 用id查询Circle
		circleRouterWithoutRecord.GET("getCircleList", circleApi.GetCircleList)                             // 分页获取Circle列表
		circleRouterWithoutRecord.GET("findCircleUser", circleApi.FindCircleUser)                           // 用id查询CircleUser
	}
	return circleRouter
}
