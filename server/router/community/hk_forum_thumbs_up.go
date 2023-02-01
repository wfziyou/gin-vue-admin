package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkForumThumbsUpRouter struct {
}

// InitHkForumThumbsUpRouter 初始化 HkForumThumbsUp 路由信息
func (s *HkForumThumbsUpRouter) InitHkForumThumbsUpRouter(Router *gin.RouterGroup) {
	hkForumThumbsUpRouter := Router.Group("hkForumThumbsUp").Use(middleware.OperationRecord())
	hkForumThumbsUpRouterWithoutRecord := Router.Group("hkForumThumbsUp")
	var hkForumThumbsUpApi = v1.ApiGroupApp.CommunityApiGroup.HkForumThumbsUpApi
	{
		hkForumThumbsUpRouter.POST("createHkForumThumbsUp", hkForumThumbsUpApi.CreateHkForumThumbsUp)   // 新建HkForumThumbsUp
		hkForumThumbsUpRouter.DELETE("deleteHkForumThumbsUp", hkForumThumbsUpApi.DeleteHkForumThumbsUp) // 删除HkForumThumbsUp
		hkForumThumbsUpRouter.DELETE("deleteHkForumThumbsUpByIds", hkForumThumbsUpApi.DeleteHkForumThumbsUpByIds) // 批量删除HkForumThumbsUp
		hkForumThumbsUpRouter.PUT("updateHkForumThumbsUp", hkForumThumbsUpApi.UpdateHkForumThumbsUp)    // 更新HkForumThumbsUp
	}
	{
		hkForumThumbsUpRouterWithoutRecord.GET("findHkForumThumbsUp", hkForumThumbsUpApi.FindHkForumThumbsUp)        // 根据ID获取HkForumThumbsUp
		hkForumThumbsUpRouterWithoutRecord.GET("getHkForumThumbsUpList", hkForumThumbsUpApi.GetHkForumThumbsUpList)  // 获取HkForumThumbsUp列表
	}
}
