package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkForumTopicGroupRouter struct {
}

// InitHkForumTopicGroupRouter 初始化 HkForumTopicGroup 路由信息
func (s *HkForumTopicGroupRouter) InitHkForumTopicGroupRouter(Router *gin.RouterGroup) {
	hkForumTopicGroupRouter := Router.Group("hkForumTopicGroup").Use(middleware.OperationRecord())
	hkForumTopicGroupRouterWithoutRecord := Router.Group("hkForumTopicGroup")
	var hkForumTopicGroupApi = v1.ApiGroupApp.CommunityApiGroup.HkForumTopicGroupApi
	{
		hkForumTopicGroupRouter.POST("createHkForumTopicGroup", hkForumTopicGroupApi.CreateHkForumTopicGroup)   // 新建HkForumTopicGroup
		hkForumTopicGroupRouter.DELETE("deleteHkForumTopicGroup", hkForumTopicGroupApi.DeleteHkForumTopicGroup) // 删除HkForumTopicGroup
		hkForumTopicGroupRouter.DELETE("deleteHkForumTopicGroupByIds", hkForumTopicGroupApi.DeleteHkForumTopicGroupByIds) // 批量删除HkForumTopicGroup
		hkForumTopicGroupRouter.PUT("updateHkForumTopicGroup", hkForumTopicGroupApi.UpdateHkForumTopicGroup)    // 更新HkForumTopicGroup
	}
	{
		hkForumTopicGroupRouterWithoutRecord.GET("findHkForumTopicGroup", hkForumTopicGroupApi.FindHkForumTopicGroup)        // 根据ID获取HkForumTopicGroup
		hkForumTopicGroupRouterWithoutRecord.GET("getHkForumTopicGroupList", hkForumTopicGroupApi.GetHkForumTopicGroupList)  // 获取HkForumTopicGroup列表
	}
}
