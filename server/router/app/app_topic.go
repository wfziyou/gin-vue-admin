package app

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TopicRouter struct {
}

func (router *TopicRouter) InitTopicRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	appRouter := Router.Group("app")
	topicRouter := appRouter.Group("topic").Use(middleware.OperationRecord())
	topicRouterWithoutRecord := appRouter.Group("topic")
	var topicApi = v1.ApiGroupApp.AppApiGroup.TopicApi
	{
		topicRouter.POST("createForumTopic", topicApi.CreateForumTopic)             //创建ForumTopic
		topicRouter.DELETE("deleteForumTopic", topicApi.DeleteForumTopic)           //删除ForumTopic
		topicRouter.DELETE("deleteForumTopicByIds", topicApi.DeleteForumTopicByIds) //批量删除ForumTopic
		topicRouter.PUT("updateForumTopic", topicApi.UpdateForumTopic)              //更新ForumTopic
	}
	{
		topicRouterWithoutRecord.GET("findForumTopicGroup", topicApi.FindForumTopicGroup)             //用id查询ForumTopicGroup
		topicRouterWithoutRecord.GET("getForumTopicGroupList", topicApi.GetForumTopicGroupList)       //分页获取ForumTopicGroup列表
		topicRouterWithoutRecord.GET("getForumTopicGroupListAll", topicApi.GetForumTopicGroupListAll) //获取ForumTopicGroup列表
		topicRouterWithoutRecord.GET("findForumTopic", topicApi.FindForumTopic)                       //用id查询ForumTopic
		topicRouterWithoutRecord.GET("getForumTopicList", topicApi.GetForumTopicList)                 //分页获取ForumTopic列表
	}
	return topicRouter
}
