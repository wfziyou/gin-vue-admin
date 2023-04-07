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
	forumPostsRouterWithoutRecord := appRouter.Group("forumPosts")
	var forumPostsApi = v1.ApiGroupApp.AppApiGroup.ForumPostsApi
	{
		forumPostsRouter.POST("createForumPosts", forumPostsApi.CreateForumPosts)                 //创建帖子
		forumPostsRouter.DELETE("deleteForumPosts", forumPostsApi.DeleteForumPosts)               //删除帖子
		forumPostsRouter.DELETE("deleteSelfForumPosts", forumPostsApi.DeleteSelfForumPosts)       //删除自己的帖子
		forumPostsRouter.DELETE("deleteForumPostsByIds", forumPostsApi.DeleteForumPostsByIds)     //批量删除帖子
		forumPostsRouter.PUT("updateForumPosts", forumPostsApi.UpdateForumPosts)                  //更新ForumPosts
		forumPostsRouter.POST("createForumComment", forumPostsApi.CreateForumComment)             //创建ForumComment
		forumPostsRouter.DELETE("deleteForumComment", forumPostsApi.DeleteForumComment)           //删除评论
		forumPostsRouter.DELETE("deleteForumCommentByIds", forumPostsApi.DeleteForumCommentByIds) //批量删除评论
		forumPostsRouter.POST("createForumThumbsUp", forumPostsApi.CreateForumThumbsUp)           //帖子点赞
		forumPostsRouter.DELETE("deleteForumThumbsUp", forumPostsApi.DeleteForumThumbsUp)         //删除帖子点赞
		forumPostsRouter.POST("createCommentThumbsUp", forumPostsApi.CreateCommentThumbsUp)       //评论点赞
		forumPostsRouter.DELETE("deleteCommentThumbsUp", forumPostsApi.DeleteCommentThumbsUp)     //删除评论点赞
	}
	{
		forumPostsRouterWithoutRecord.GET("findForumPosts", forumPostsApi.FindForumPosts)           //用id查询帖子
		forumPostsRouterWithoutRecord.GET("getForumPostsList", forumPostsApi.GetForumPostsList)     //分页获取ForumPosts列表
		forumPostsRouterWithoutRecord.GET("findForumComment", forumPostsApi.FindForumComment)       //用id查询ForumComment
		forumPostsRouterWithoutRecord.GET("getForumCommentList", forumPostsApi.GetForumCommentList) //分页获取评论列表

	}
	return forumPostsRouter
}
