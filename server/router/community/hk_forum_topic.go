package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkForumTopicRouter struct {
}

// InitHkForumTopicRouter 初始化 HkForumTopic 路由信息
func (s *HkForumTopicRouter) InitHkForumTopicRouter(Router *gin.RouterGroup) {
	hkForumTopicRouter := Router.Group("hkForumTopic").Use(middleware.OperationRecord())
	hkForumTopicRouterWithoutRecord := Router.Group("hkForumTopic")
	var hkForumTopicApi = v1.ApiGroupApp.CommunityApiGroup.HkForumTopicApi
	{
		hkForumTopicRouter.POST("createHkForumTopic", hkForumTopicApi.CreateHkForumTopic)   // 新建HkForumTopic
		hkForumTopicRouter.DELETE("deleteHkForumTopic", hkForumTopicApi.DeleteHkForumTopic) // 删除HkForumTopic
		hkForumTopicRouter.DELETE("deleteHkForumTopicByIds", hkForumTopicApi.DeleteHkForumTopicByIds) // 批量删除HkForumTopic
		hkForumTopicRouter.PUT("updateHkForumTopic", hkForumTopicApi.UpdateHkForumTopic)    // 更新HkForumTopic
	}
	{
		hkForumTopicRouterWithoutRecord.GET("findHkForumTopic", hkForumTopicApi.FindHkForumTopic)        // 根据ID获取HkForumTopic
		hkForumTopicRouterWithoutRecord.GET("getHkForumTopicList", hkForumTopicApi.GetHkForumTopicList)  // 获取HkForumTopic列表
	}
}
