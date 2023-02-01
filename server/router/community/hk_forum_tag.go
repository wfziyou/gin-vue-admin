package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkForumTagRouter struct {
}

// InitHkForumTagRouter 初始化 HkForumTag 路由信息
func (s *HkForumTagRouter) InitHkForumTagRouter(Router *gin.RouterGroup) {
	hkForumTagRouter := Router.Group("hkForumTag").Use(middleware.OperationRecord())
	hkForumTagRouterWithoutRecord := Router.Group("hkForumTag")
	var hkForumTagApi = v1.ApiGroupApp.CommunityApiGroup.HkForumTagApi
	{
		hkForumTagRouter.POST("createHkForumTag", hkForumTagApi.CreateHkForumTag)   // 新建HkForumTag
		hkForumTagRouter.DELETE("deleteHkForumTag", hkForumTagApi.DeleteHkForumTag) // 删除HkForumTag
		hkForumTagRouter.DELETE("deleteHkForumTagByIds", hkForumTagApi.DeleteHkForumTagByIds) // 批量删除HkForumTag
		hkForumTagRouter.PUT("updateHkForumTag", hkForumTagApi.UpdateHkForumTag)    // 更新HkForumTag
	}
	{
		hkForumTagRouterWithoutRecord.GET("findHkForumTag", hkForumTagApi.FindHkForumTag)        // 根据ID获取HkForumTag
		hkForumTagRouterWithoutRecord.GET("getHkForumTagList", hkForumTagApi.GetHkForumTagList)  // 获取HkForumTag列表
	}
}
