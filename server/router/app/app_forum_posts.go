package app

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ForumPostsRouter struct {
}

func (router *ForumPostsRouter) InitForumPostsRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	appRouter := Router.Group("app")
	forumPostsRouter := appRouter.Group("forumPosts").Use(middleware.OperationRecord())
	userRouterWithoutRecord := appRouter.Group("forumPosts")
	var api = v1.ApiGroupApp.AppApiGroup.ForumPostsApi
	{
		forumPostsRouter.POST("createForumPosts", api.CreateForumPosts)             //创建帖子
		forumPostsRouter.POST("createNews", api.CreateNews)                         //创建资讯
		forumPostsRouter.DELETE("deleteForumPosts", api.DeleteForumPosts)           //删除帖子
		forumPostsRouter.DELETE("deleteForumPostsByIds", api.DeleteForumPostsByIds) //批量删除帖子
		forumPostsRouter.DELETE("deleteSelfForumPosts", api.DeleteSelfForumPosts)   //删除自己的帖子

		forumPostsRouter.PUT("updateForumPosts", api.UpdateForumPosts)                  //更新帖子
		forumPostsRouter.POST("createForumComment", api.CreateForumComment)             //创建评论
		forumPostsRouter.DELETE("deleteForumComment", api.DeleteForumComment)           //删除评论
		forumPostsRouter.DELETE("deleteForumCommentByIds", api.DeleteForumCommentByIds) //批量删除评论
		forumPostsRouter.POST("createForumThumbsUp", api.CreateForumThumbsUp)           //帖子点赞
		forumPostsRouter.DELETE("deleteForumThumbsUp", api.DeleteForumThumbsUp)         //删除帖子点赞
		forumPostsRouter.POST("createCommentThumbsUp", api.CreateCommentThumbsUp)       //评论点赞
		forumPostsRouter.DELETE("deleteCommentThumbsUp", api.DeleteCommentThumbsUp)     //删除评论点赞
	}
	{
		userRouterWithoutRecord.GET("getRecommendPostsList", api.GetRecommendPostsList)             //分页获取推荐帖子列表
		userRouterWithoutRecord.GET("getGlobalTopInfoList", api.GetGlobalTopInfoList)               //获全局置顶资讯列表
		userRouterWithoutRecord.GET("getGlobalRecommendInfoList", api.GetGlobalRecommendInfoList)   //分页获全局推荐资讯列表
		userRouterWithoutRecord.GET("getNearbyRecommendPostsList", api.GetNearbyRecommendPostsList) //分页获附近推荐帖子列表

		userRouterWithoutRecord.GET("findForumPosts", api.FindForumPosts)               //用id查询帖子
		userRouterWithoutRecord.GET("getForumPostsList", api.GetForumPostsList)         //分页获取帖子列表
		userRouterWithoutRecord.GET("findForumComment", api.FindForumComment)           //用id查询评论
		userRouterWithoutRecord.GET("getForumCommentList", api.GetForumCommentList)     //分页获取评论列表
		userRouterWithoutRecord.GET("getUserForumPostsList", api.GetUserForumPostsList) // 分页获取用户帖子列表

	}
	return forumPostsRouter
}
