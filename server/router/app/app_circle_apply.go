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
		circleApplyRouterWithoutRecord.GET("findApply", circleApplyApi.FindApply)                                   //用id查询Apply
		circleApplyRouterWithoutRecord.GET("getApplyList", circleApplyApi.GetApplyList)                             //分页获取Apply列表
		circleApplyRouterWithoutRecord.GET("getApplyListAll", circleApplyApi.GetApplyListAll)                       //获取Apply列表
		circleApplyRouterWithoutRecord.GET("findCircleApply", circleApplyApi.FindCircleApply)                       //用id查询CircleApply
		circleApplyRouterWithoutRecord.GET("getCircleApplyList", circleApplyApi.GetCircleApplyList)                 //分页获取CircleApply列表
		circleApplyRouterWithoutRecord.GET("getCircleApplyListAll", circleApplyApi.GetCircleApplyListAll)           //获取CircleApply列表
		circleApplyRouterWithoutRecord.GET("getCircleApplyGroupList", circleApplyApi.GetCircleApplyGroupList)       //分页获取CircleApplyGroup列表
		circleApplyRouterWithoutRecord.GET("getCircleApplyGroupListAll", circleApplyApi.GetCircleApplyGroupListAll) //获取CircleApplyGroup列表
		circleApplyRouterWithoutRecord.GET("getCircleClassifyListAll", circleApplyApi.GetCircleApplyGroupListAll)   //获取CircleApplyGroup列表
		circleApplyRouterWithoutRecord.GET("getCircleHotApplyList", circleApplyApi.GetCircleHotApplyList)           //获取圈子热门应用列表
	}
	return circleApplyRouter
}
