package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkForumTagGroupRouter struct {
}

// InitHkForumTagGroupRouter 初始化 HkForumTagGroup 路由信息
func (s *HkForumTagGroupRouter) InitHkForumTagGroupRouter(Router *gin.RouterGroup) {
	hkForumTagGroupRouter := Router.Group("hkForumTagGroup").Use(middleware.OperationRecord())
	hkForumTagGroupRouterWithoutRecord := Router.Group("hkForumTagGroup")
	var hkForumTagGroupApi = v1.ApiGroupApp.CommunityApiGroup.HkForumTagGroupApi
	{
		hkForumTagGroupRouter.POST("createHkForumTagGroup", hkForumTagGroupApi.CreateHkForumTagGroup)   // 新建HkForumTagGroup
		hkForumTagGroupRouter.DELETE("deleteHkForumTagGroup", hkForumTagGroupApi.DeleteHkForumTagGroup) // 删除HkForumTagGroup
		hkForumTagGroupRouter.DELETE("deleteHkForumTagGroupByIds", hkForumTagGroupApi.DeleteHkForumTagGroupByIds) // 批量删除HkForumTagGroup
		hkForumTagGroupRouter.PUT("updateHkForumTagGroup", hkForumTagGroupApi.UpdateHkForumTagGroup)    // 更新HkForumTagGroup
	}
	{
		hkForumTagGroupRouterWithoutRecord.GET("findHkForumTagGroup", hkForumTagGroupApi.FindHkForumTagGroup)        // 根据ID获取HkForumTagGroup
		hkForumTagGroupRouterWithoutRecord.GET("getHkForumTagGroupList", hkForumTagGroupApi.GetHkForumTagGroupList)  // 获取HkForumTagGroup列表
	}
}
