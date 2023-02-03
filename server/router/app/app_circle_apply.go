package app

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CircleApplyRouter struct {
}

func (router *CircleApplyRouter) InitCircleApplyRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	appRouter := Router.Group("app")
	circleApplyRouter := appRouter.Group("circleApply").Use(middleware.OperationRecord())
	circleApplyRouterWithoutRecord := appRouter.Group("circleApply")
	var circleApplyApi = v1.ApiGroupApp.AppApiGroup.CircleApplyApi
	{
		circleApplyRouter.POST("createUserCircleApply", circleApplyApi.CreateUserCircleApply)             //创建UserCircleApply
		circleApplyRouter.DELETE("deleteUserCircleApply", circleApplyApi.DeleteUserCircleApply)           //删除UserCircleApply
		circleApplyRouter.DELETE("deleteUserCircleApplyByIds", circleApplyApi.DeleteUserCircleApplyByIds) //批量删除UserCircleApply
	}
	{
		circleApplyRouterWithoutRecord.GET("getUserCircleApplyListALL", circleApplyApi.GetUserCircleApplyListALL)   //获取UserCircleApply列表
		circleApplyRouterWithoutRecord.GET("findApply", circleApplyApi.FindApply)                                   //用id查询Apply
		circleApplyRouterWithoutRecord.GET("getApplyList", circleApplyApi.GetApplyList)                             //分页获取Apply列表
		circleApplyRouterWithoutRecord.GET("getApplyListAll", circleApplyApi.GetApplyListAll)                       //获取Apply列表
		circleApplyRouterWithoutRecord.GET("findCircleApply", circleApplyApi.FindCircleApply)                       //用id查询CircleApply
		circleApplyRouterWithoutRecord.GET("getCircleApplyList", circleApplyApi.GetCircleApplyList)                 //分页获取CircleApply列表
		circleApplyRouterWithoutRecord.GET("getCircleApplyListAll", circleApplyApi.GetCircleApplyListAll)           //获取CircleApply列表
		circleApplyRouterWithoutRecord.GET("getCircleApplyGroupList", circleApplyApi.GetCircleApplyGroupList)       //分页获取CircleApplyGroup列表
		circleApplyRouterWithoutRecord.GET("getCircleApplyGroupListAll", circleApplyApi.GetCircleApplyGroupListAll) //获取CircleApplyGroup列表
	}
	return circleApplyRouter
}
