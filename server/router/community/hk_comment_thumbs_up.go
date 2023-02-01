package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkCommentThumbsUpRouter struct {
}

// InitHkCommentThumbsUpRouter 初始化 HkCommentThumbsUp 路由信息
func (s *HkCommentThumbsUpRouter) InitHkCommentThumbsUpRouter(Router *gin.RouterGroup) {
	hkCommentThumbsUpRouter := Router.Group("hkCommentThumbsUp").Use(middleware.OperationRecord())
	hkCommentThumbsUpRouterWithoutRecord := Router.Group("hkCommentThumbsUp")
	var hkCommentThumbsUpApi = v1.ApiGroupApp.CommunityApiGroup.HkCommentThumbsUpApi
	{
		hkCommentThumbsUpRouter.POST("createHkCommentThumbsUp", hkCommentThumbsUpApi.CreateHkCommentThumbsUp)   // 新建HkCommentThumbsUp
		hkCommentThumbsUpRouter.DELETE("deleteHkCommentThumbsUp", hkCommentThumbsUpApi.DeleteHkCommentThumbsUp) // 删除HkCommentThumbsUp
		hkCommentThumbsUpRouter.DELETE("deleteHkCommentThumbsUpByIds", hkCommentThumbsUpApi.DeleteHkCommentThumbsUpByIds) // 批量删除HkCommentThumbsUp
		hkCommentThumbsUpRouter.PUT("updateHkCommentThumbsUp", hkCommentThumbsUpApi.UpdateHkCommentThumbsUp)    // 更新HkCommentThumbsUp
	}
	{
		hkCommentThumbsUpRouterWithoutRecord.GET("findHkCommentThumbsUp", hkCommentThumbsUpApi.FindHkCommentThumbsUp)        // 根据ID获取HkCommentThumbsUp
		hkCommentThumbsUpRouterWithoutRecord.GET("getHkCommentThumbsUpList", hkCommentThumbsUpApi.GetHkCommentThumbsUpList)  // 获取HkCommentThumbsUp列表
	}
}
