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
		circleApplyRouterWithoutRecord.GET("getCircleHotApplyList", circleApplyApi.GetCircleHotApplyList)           //获取圈子热门应用列表
		circleApplyRouter.POST("setCircleHotApply", circleApplyApi.SetCircleHotApply)                               //(圈子管理者)设置圈子热门应用
		circleApplyRouter.POST("addCircleApply", circleApplyApi.AddCircleApply)                                     //(圈子管理者)添加圈子应用
		circleApplyRouter.POST("updateCircleApply", circleApplyApi.UpdateCircleApply)                               //(圈子管理者)更新圈子应用
		circleApplyRouter.DELETE("deleteCircleApply", circleApplyApi.DeleteCircleApply)                             //(圈子管理者)删除圈子应用
		circleApplyRouter.POST("addCircleApplyGroup", circleApplyApi.AddCircleApplyGroup)                           //(圈子管理者)添加圈子应用分组
		circleApplyRouter.DELETE("deleteCircleApplyGroup", circleApplyApi.DeleteCircleApplyGroup)                   //(圈子管理者)删除圈子应用分组
		circleApplyRouter.POST("setCircleApplyGroupSort", circleApplyApi.SetCircleApplyGroupSort)                   //(圈子管理者)设置圈子应用分组顺序
	}
	return circleApplyRouter
}
