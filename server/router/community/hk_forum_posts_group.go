package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkForumPostsGroupRouter struct {
}

// InitHkForumPostsGroupRouter 初始化 HkForumPostsGroup 路由信息
func (s *HkForumPostsGroupRouter) InitHkForumPostsGroupRouter(Router *gin.RouterGroup) {
	hkForumPostsGroupRouter := Router.Group("hkForumPostsGroup").Use(middleware.OperationRecord())
	hkForumPostsGroupRouterWithoutRecord := Router.Group("hkForumPostsGroup")
	var hkForumPostsGroupApi = v1.ApiGroupApp.CommunityApiGroup.HkForumPostsGroupApi
	{
		hkForumPostsGroupRouter.POST("createHkForumPostsGroup", hkForumPostsGroupApi.CreateHkForumPostsGroup)   // 新建HkForumPostsGroup
		hkForumPostsGroupRouter.DELETE("deleteHkForumPostsGroup", hkForumPostsGroupApi.DeleteHkForumPostsGroup) // 删除HkForumPostsGroup
		hkForumPostsGroupRouter.DELETE("deleteHkForumPostsGroupByIds", hkForumPostsGroupApi.DeleteHkForumPostsGroupByIds) // 批量删除HkForumPostsGroup
		hkForumPostsGroupRouter.PUT("updateHkForumPostsGroup", hkForumPostsGroupApi.UpdateHkForumPostsGroup)    // 更新HkForumPostsGroup
	}
	{
		hkForumPostsGroupRouterWithoutRecord.GET("findHkForumPostsGroup", hkForumPostsGroupApi.FindHkForumPostsGroup)        // 根据ID获取HkForumPostsGroup
		hkForumPostsGroupRouterWithoutRecord.GET("getHkForumPostsGroupList", hkForumPostsGroupApi.GetHkForumPostsGroupList)  // 获取HkForumPostsGroup列表
	}
}
