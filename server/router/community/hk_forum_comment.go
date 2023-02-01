package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkForumCommentRouter struct {
}

// InitHkForumCommentRouter 初始化 HkForumComment 路由信息
func (s *HkForumCommentRouter) InitHkForumCommentRouter(Router *gin.RouterGroup) {
	hkForumCommentRouter := Router.Group("hkForumComment").Use(middleware.OperationRecord())
	hkForumCommentRouterWithoutRecord := Router.Group("hkForumComment")
	var hkForumCommentApi = v1.ApiGroupApp.CommunityApiGroup.HkForumCommentApi
	{
		hkForumCommentRouter.POST("createHkForumComment", hkForumCommentApi.CreateHkForumComment)   // 新建HkForumComment
		hkForumCommentRouter.DELETE("deleteHkForumComment", hkForumCommentApi.DeleteHkForumComment) // 删除HkForumComment
		hkForumCommentRouter.DELETE("deleteHkForumCommentByIds", hkForumCommentApi.DeleteHkForumCommentByIds) // 批量删除HkForumComment
		hkForumCommentRouter.PUT("updateHkForumComment", hkForumCommentApi.UpdateHkForumComment)    // 更新HkForumComment
	}
	{
		hkForumCommentRouterWithoutRecord.GET("findHkForumComment", hkForumCommentApi.FindHkForumComment)        // 根据ID获取HkForumComment
		hkForumCommentRouterWithoutRecord.GET("getHkForumCommentList", hkForumCommentApi.GetHkForumCommentList)  // 获取HkForumComment列表
	}
}
