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
		circleRouter.POST("createCircleRequest", circleApi.CreateCircleRequest) //创建圈子请求
		circleRouter.PUT("updateCircle", circleApi.UpdateCircle)                //(圈子管理者)更新圈子
		circleRouter.POST("setUserCurCircle", circleApi.SetUserCurCircle)       //设置用户当前圈子
		circleRouter.POST("enterCircle", circleApi.EnterCircle)                 //加入圈子
		circleRouter.POST("exitCircle", circleApi.ExitCircle)                   //退出圈子
		circleRouter.POST("applyEnterCircle", circleApi.ApplyEnterCircle)       //申请加入圈子

		circleRouter.DELETE("deleteCircleUser", circleApi.DeleteCircleUser)   //删除圈子用户
		circleRouter.DELETE("deleteCircleUsers", circleApi.DeleteCircleUsers) //批量删除圈子用户
		circleRouter.PUT("updateCircleUser", circleApi.UpdateCircleUser)      //更新圈子用户
	}
	{
		circleRouterWithoutRecord.GET("getCircleForumPostsList", circleApi.GetCircleForumPostsList)         // 分页获取圈子的帖子列表
		circleRouterWithoutRecord.GET("getUserCircleForumPostsList", circleApi.GetUserCircleForumPostsList) // 分页获取用户圈子ForumPosts列表
		circleRouterWithoutRecord.GET("getSelfCircleList", circleApi.GetSelfCircleList)                     // 分页获取用户加入的圈子列表
		circleRouterWithoutRecord.GET("findCircle", circleApi.FindCircle)                                   // 用id查询Circle
		circleRouterWithoutRecord.GET("getCircleList", circleApi.GetCircleList)                             // 分页获取圈子列表
		circleRouterWithoutRecord.GET("findCircleUser", circleApi.FindCircleUser)                           // 用id查询CircleUser
		circleRouterWithoutRecord.GET("getCircleUserList", circleApi.GetCircleUserList)                     // 分页获取圈子的用户信息列表

		circleRouterWithoutRecord.GET("getCircleClassifyList", circleApi.GetCircleClassifyList)         // 分页获取圈子分类列表
		circleRouterWithoutRecord.GET("getCircleClassifyListAll", circleApi.GetCircleClassifyListAll)   // 获取圈子分类列表
		circleRouterWithoutRecord.GET("enterCircleApplyList", circleApi.EnterCircleApplyList)           // 分页获取加入圈子申请
		circleRouterWithoutRecord.GET("approveEnterCircleRequest", circleApi.ApproveEnterCircleRequest) // 审批加入圈子申请

	}
	return circleRouter
}
