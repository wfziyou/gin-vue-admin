package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type HkForumPostsRouter struct {
}

// InitHkForumPostsRouter 初始化 HkForumPosts 路由信息
func (s *HkForumPostsRouter) InitHkForumPostsRouter(Router *gin.RouterGroup) {
	hkForumPostsRouter := Router.Group("hkForumPosts").Use(middleware.OperationRecord())
	hkForumPostsRouterWithoutRecord := Router.Group("hkForumPosts")
	var hkForumPostsApi = v1.ApiGroupApp.CommunityApiGroup.HkForumPostsApi
	{
		hkForumPostsRouter.POST("createHkForumPosts", hkForumPostsApi.CreateHkForumPosts)   // 新建HkForumPosts
		hkForumPostsRouter.DELETE("deleteHkForumPosts", hkForumPostsApi.DeleteHkForumPosts) // 删除HkForumPosts
		hkForumPostsRouter.DELETE("deleteHkForumPostsByIds", hkForumPostsApi.DeleteHkForumPostsByIds) // 批量删除HkForumPosts
		hkForumPostsRouter.PUT("updateHkForumPosts", hkForumPostsApi.UpdateHkForumPosts)    // 更新HkForumPosts
	}
	{
		hkForumPostsRouterWithoutRecord.GET("findHkForumPosts", hkForumPostsApi.FindHkForumPosts)        // 根据ID获取HkForumPosts
		hkForumPostsRouterWithoutRecord.GET("getHkForumPostsList", hkForumPostsApi.GetHkForumPostsList)  // 获取HkForumPosts列表
	}
}
